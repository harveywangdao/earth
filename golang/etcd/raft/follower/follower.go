package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/harveywangdao/earth/golang/etcd/raft/common"
)

func joincluster(leaderAddr, addr string) {
	time.Sleep(time.Second * 8)
	cli := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s/v1/api/node/%s", leaderAddr, addr), nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	resp, err := cli.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("join cluster fail")
	}
}

func follower(addr, dataDir, leaderAddr string) error {
	raftNode, ch, err := common.StartRaftServer(addr, dataDir)
	if err != nil {
		log.Println(err)
		return err
	}

	go joincluster(leaderAddr, addr)

	writetk := time.NewTicker(time.Second * 5)
	defer writetk.Stop()
	readtk := time.NewTicker(time.Second * 6)
	defer readtk.Stop()
	count := 0

	for {
		select {
		case <-ch:
			log.Println("leadership changed")

		case <-writetk.C:
			event := common.LogEntryData{Key: fmt.Sprintf("%s--key--%d", dataDir, count), Value: fmt.Sprintf("%s--value--%d", dataDir, count)}
			eventBytes, err := json.Marshal(event)
			if err != nil {
				log.Println(err)
				return err
			}
			future := raftNode.Apply(eventBytes, 5*time.Second)
			if err := future.Error(); err != nil {
				log.Println(err)
			}
			count++
		case <-readtk.C:
		}
	}

	return nil
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	leader := flag.String("leader", "192.168.126.128:6661", "leader")
	addr := flag.String("addr", "192.168.126.128:7772", "addr")
	dataDir := flag.String("dataDir", "follower1", "dataDir")
	flag.Parse()

	follower(*addr, *dataDir, *leader)
}
