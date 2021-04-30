package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"
)

var (
	addNodeCh = make(chan string)
)

type snapshot struct {
}

func (s *snapshot) Persist(sink raft.SnapshotSink) error {
	log.Println("Persist")
	if _, err := sink.Write([]byte("sss")); err != nil {
		sink.Cancel()
		return err
	}

	if err := sink.Close(); err != nil {
		sink.Cancel()
		return err
	}
	return nil
}

func (s *snapshot) Release() {
	log.Println("Release")
}

type FSM struct {
	log *log.Logger
}

type logEntryData struct {
	Key   string
	Value string
}

func (f *FSM) Apply(logEntry *raft.Log) interface{} {
	e := logEntryData{}
	if err := json.Unmarshal(logEntry.Data, &e); err != nil {
		panic("Failed unmarshaling Raft log entry. This is a bug.")
	}
	log.Println(e)
	return e.Key + e.Value
}

func (f *FSM) Snapshot() (raft.FSMSnapshot, error) {
	log.Println("Snapshot")
	return &snapshot{}, nil
}

func (f *FSM) Restore(rc io.ReadCloser) error {
	buf := make([]byte, 32)
	rc.Read(buf)
	defer rc.Close()
	log.Println("Restore", string(buf))
	return nil
}

type _log struct{}

func (l *_log) Write(p []byte) (n int, err error) {
	log.Println(string(p))
	return len(p), nil
}

func common(addr, dataDir string) (*raft.Raft, chan bool, error) {
	ch := make(chan bool, 1)
	raftConfig := raft.DefaultConfig()
	raftConfig.LocalID = raft.ServerID(addr)
	raftConfig.Logger = hclog.New(&hclog.LoggerOptions{
		Name:   "raft: " + dataDir + " ",
		Output: &_log{},
	})
	raftConfig.SnapshotInterval = 20 * time.Second
	raftConfig.SnapshotThreshold = 2
	raftConfig.NotifyCh = ch

	address, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}
	transport, err := raft.NewTCPTransport(address.String(), address, 3, 10*time.Second, os.Stderr)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}
	if err := os.RemoveAll(dataDir); err != nil {
		log.Fatal(err)
		return nil, nil, err
	}
	if err := os.MkdirAll(dataDir, 0700); err != nil {
		log.Fatal(err)
		return nil, nil, err
	}
	fsm := &FSM{
		log: log.New(os.Stderr, "FSM: ", log.Ldate|log.Ltime),
	}
	snapshotStore, err := raft.NewFileSnapshotStore(dataDir, 1, os.Stderr)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}
	logStore, err := raftboltdb.NewBoltStore(filepath.Join(dataDir, "raft-log.bolt"))
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}
	stableStore, err := raftboltdb.NewBoltStore(filepath.Join(dataDir, "raft-stable.bolt"))
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}
	raftNode, err := raft.NewRaft(raftConfig, fsm, logStore, stableStore, snapshotStore, transport)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}

	future := raftNode.BootstrapCluster(raft.Configuration{
		Servers: []raft.Server{
			{
				ID:      raftConfig.LocalID,
				Address: transport.LocalAddr(),
			},
		},
	})
	if err := future.Error(); err != nil {
		log.Println(dataDir, err)
	}
	return raftNode, ch, nil
}

func do1(addr, dataDir string) error {
	raftNode, ch, err := common(addr, dataDir)
	if err != nil {
		log.Println(err)
		return err
	}

	writetk := time.NewTicker(time.Second * 5)
	defer writetk.Stop()
	readtk := time.NewTicker(time.Second * 6)
	defer readtk.Stop()
	c := 0

	for {
		select {
		case <-ch:
			log.Println(dataDir, addr)

			for i := 0; i < 5; i++ {
				event := logEntryData{Key: "key", Value: fmt.Sprintf("%s--%d", dataDir, i)}
				eventBytes, err := json.Marshal(event)
				if err != nil {
					log.Println(dataDir, err)
					return err
				}
				future := raftNode.Apply(eventBytes, 5*time.Second)
				if err := future.Error(); err != nil {
					log.Println(dataDir, err)
				}
			}

		case newAddr := <-addNodeCh:
			log.Println(dataDir, "add node:", newAddr)
			addPeerFuture := raftNode.AddVoter(raft.ServerID(newAddr), raft.ServerAddress(newAddr), 0, 0)
			if err := addPeerFuture.Error(); err != nil {
				log.Println(dataDir, err)
			}

		case <-writetk.C:
			/*event := logEntryData{Key: "key", Value: fmt.Sprintf("%s--%d", dataDir, c)}
			eventBytes, err := json.Marshal(event)
			if err != nil {
				log.Println(dataDir, err)
				return err
			}
			future := raftNode.Apply(eventBytes, 5*time.Second)
			if err := future.Error(); err != nil {
				log.Println(dataDir, err)
			}
			*/
			c++
		case <-readtk.C:
		}
	}

	return nil
}

func do2(addr, dataDir string) error {
	raftNode, ch, err := common(addr, dataDir)
	if err != nil {
		log.Println(err)
		return err
	}
	_ = raftNode

	addNodeCh <- addr
	readtk := time.NewTicker(time.Second * 6)
	defer readtk.Stop()

	for {
		select {
		case <-ch:
			log.Println(dataDir, addr)
		case <-readtk.C:
			/*event := logEntryData{Key: "key", Value: dataDir}
			eventBytes, err := json.Marshal(event)
			if err != nil {
				log.Println(dataDir, err)
				return err
			}
			future := raftNode.Apply(eventBytes, 5*time.Second)
			if err := future.Error(); err != nil {
				log.Println(dataDir, err)
			}*/
		}
	}

	return nil
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	go do1("192.168.126.128:7771", "node1")

	time.Sleep(time.Second * 5)
	go do2("192.168.126.128:7772", "node2")

	time.Sleep(time.Second * 5)
	go do2("192.168.126.128:7773", "node3")

	select {}
}
