package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/harveywangdao/earth/golang/etcd/raft/common"
	"github.com/hashicorp/raft"
)

var (
	addNodeCh = make(chan string)
)

func leader(addr, dataDir string) error {
	raftNode, ch, err := common.StartRaftServer(addr, dataDir)
	if err != nil {
		log.Println(err)
		return err
	}

	writetk := time.NewTicker(time.Second * 5)
	defer writetk.Stop()
	readtk := time.NewTicker(time.Second * 6)
	defer readtk.Stop()
	count := 0

	for {
		select {
		case <-ch:
			log.Println("leadership changed")

		case newAddr := <-addNodeCh:
			log.Println("add node:", newAddr)
			addPeerFuture := raftNode.AddVoter(raft.ServerID(newAddr), raft.ServerAddress(newAddr), 0, 0)
			if err := addPeerFuture.Error(); err != nil {
				log.Println(err)
			}

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

func httpserver(ipport string) {
	router := gin.Default()
	router.POST("/v1/api/node/:addr", func(c *gin.Context) {
		addr := c.Param("addr")
		if addr != "" {
			addNodeCh <- addr
		}
	})
	log.Fatal(router.Run(ipport))
}

/*
1.从节点加入之前,主节点增加自己的数据,从节点也增加自己的数据
2.从节点加入之后,数据如何同步
3.主节点宕机,从节点如何选举

选举:
term
log version
commit index
apply
snapshot

Leader
Follower
Candidate

*/
func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	addr := flag.String("addr", "192.168.126.128:7771", "addr")
	dataDir := flag.String("dataDir", "leader", "dataDir")
	httpport := flag.String("httpport", "192.168.126.128:6661", "httpport")
	flag.Parse()

	go httpserver(*httpport)

	leader(*addr, *dataDir)
}
