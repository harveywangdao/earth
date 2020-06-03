package main

import (
	"log"
	"time"

	"github.com/bsm/redislock"
	"github.com/go-redis/redis"
)

/*
1.正常步骤
2.key已经存在  会阻塞,在超时前获取到key set redis-dis-lock aaaaa ex 20
3.key已经存在  会阻塞,在超时前也没获取到key set redis-dis-lock aaaaa ex 60
4.key超时,但任务还没执行完
5.key超时,但任务还没执行完,key又被另一个任务获取到
6.key超时,但任务还没执行完,key又被另一个任务获取到,用保活阻止这一情况

注意:由于网络不确定因素,可能会保活失败,会出现同一把锁被两个节点拿去
*/

func do1() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	log.Println(pong, err)

	option := &redislock.Options{
		RetryStrategy: redislock.LinearBackoff(100 * time.Millisecond), // 锁失败后的重试间隔时间
	}
	locker := redislock.New(client)
	lock, err := locker.Obtain("redis-dis-lock", 10*time.Second, option) // redis key过期时间
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("get key")

	stopTimerCh := make(chan bool)
	timer := time.NewTimer(9 * time.Second)
	go func() {
		for {
			select {
			case <-timer.C:
				if err := lock.Refresh(10*time.Second, nil); err != nil {
					log.Println(err)
					continue
				}
				log.Println("refresh success")
				timer.Reset(9 * time.Second)
			case <-stopTimerCh:
				log.Println("stopTimerCh")
				return
			}
		}
	}()

	go do2(client)
	time.Sleep(100 * time.Second)

	timer.Stop()
	stopTimerCh <- true
	if err := lock.Release(); err != nil {
		log.Println(err)
		return
	}
	log.Println("del key")
}

func do2(client *redis.Client) {
	option := &redislock.Options{
		RetryStrategy: redislock.LinearBackoff(100 * time.Millisecond), // 锁失败后的重试间隔时间
	}
	locker := redislock.New(client)
	lock, err := locker.Obtain("redis-dis-lock", 30*time.Second, option) // redis key过期时间
	if err != nil {
		log.Println("do2:", err)
		return
	}
	log.Println("do2: get key")

	time.Sleep(25 * time.Second)

	if err := lock.Release(); err != nil {
		log.Println("do2:", err)
		return
	}
	log.Println("do2: del key")
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do1()
	time.Sleep(1 * time.Hour)
}
