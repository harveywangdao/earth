package common

import (
	"encoding/json"
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

type snapshot struct {
}

func (s *snapshot) Persist(sink raft.SnapshotSink) error {
	log.Printf("Persist: id: %s\n", sink.ID())
	if _, err := sink.Write([]byte("sss")); err != nil {
		log.Println(err)
		sink.Cancel()
		return err
	}
	if err := sink.Close(); err != nil {
		log.Println(err)
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

type LogEntryData struct {
	Key   string
	Value string
}

func (f *FSM) Apply(logEntry *raft.Log) interface{} {
	entry := LogEntryData{}
	if err := json.Unmarshal(logEntry.Data, &entry); err != nil {
		log.Println(err)
		return nil
	}
	log.Println(entry)
	return entry.Key + entry.Value
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

func StartRaftServer(addr, dataDir string) (*raft.Raft, chan bool, error) {
	ch := make(chan bool, 1)

	raftConfig := raft.DefaultConfig()
	raftConfig.LocalID = raft.ServerID(addr)
	raftConfig.Logger = hclog.New(&hclog.LoggerOptions{
		Name:   "raft: ",
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
		log.Println(err)
	}
	return raftNode, ch, nil
}
