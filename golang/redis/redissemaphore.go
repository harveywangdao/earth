package main

import (
	"encoding/hex"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/satori/go.uuid"
	"log"
	"sync"
	"time"
)

type RedisSemaphore struct {
	Client      *redis.Client
	SemKey      string
	SemMember   string
	MaxSemCount int64
	SemTimeout  time.Duration // 如果节点挂掉无法释放信号量,信号量最多存活的时间
}

func (r *RedisSemaphore) TryWait() (bool, error) {
	// 清除过期的信号
	// 获取信号数量,如果小于最大信号量就增加一个信号量,否则返回失败
	script := redis.NewScript(`
    redis.call('zremrangebyscore', KEYS[1], '-inf', ARGV[1])
    if redis.call('zcard', KEYS[1]) < tonumber(ARGV[2]) then
      return redis.call('zadd', KEYS[1], ARGV[3], ARGV[4])
    else
      return -1
    end`)

	start := time.Now()
	defer func() {
		log.Printf("get redis semaphore cost %d ms\n", time.Now().Sub(start).Nanoseconds()/1000000)
	}()

	n, err := script.Run(r.Client, []string{r.SemKey}, time.Now().Add(-1*r.SemTimeout).UnixNano(), r.MaxSemCount, time.Now().UnixNano(), r.SemMember).Int64()
	if err != nil {
		return false, err
	} else if n < 0 {
		return false, nil
	}

	return true, nil
}

func (r *RedisSemaphore) Refresh() error {
	// 先判断member是否存在,如果存在就更新score,不存在就返回错
	// 正确情况返回0
	// 如果member不存在返回-1
	script := redis.NewScript(`
	    if redis.call('zscore', KEYS[1], ARGV[1]) then
	      return redis.call('zadd', KEYS[1], ARGV[2], ARGV[1])
	    else
	      return -1
	    end`)

	start := time.Now()
	defer func() {
		log.Printf("refresh redis semaphore cost %d ms\n", time.Now().Sub(start).Nanoseconds()/1000000)
	}()

	n, err := script.Run(r.Client, []string{r.SemKey}, r.SemMember, time.Now().UnixNano()).Int64()
	if err != nil {
		return err
	} else if n != 0 {
		return fmt.Errorf("refresh fail, maybe member not existed")
	}

	return nil
}

func (r *RedisSemaphore) Release() error {
	n, err := r.Client.ZRem(r.SemKey, r.SemMember).Result()
	if err != nil {
		return err
	} else if n != 1 {
		return fmt.Errorf("sem: %s member: %s not existed", r.SemKey, r.SemMember)
	}
	return nil
}

/*
1.正常流程下,如果信号量过大会不会影响性能
*/

func do1(client *redis.Client, semMember string) {
	sem := &RedisSemaphore{
		Client:      client,
		SemKey:      "redis-sem",
		SemMember:   semMember,
		MaxSemCount: 10,
		SemTimeout:  10 * time.Second,
	}

	success, err := sem.TryWait()
	if err != nil {
		log.Println(err)
		return
	} else if !success {
		log.Printf("SemKey: %s exceed max sem count: %d\n", sem.SemKey, sem.MaxSemCount)
		return
	}
	log.Println("get sem success")

	stopTimerCh := make(chan bool)
	timer := time.NewTimer(9 * time.Second)
	go func() {
		for {
			select {
			case <-timer.C:
				if err := sem.Refresh(); err != nil {
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

	time.Sleep(30 * time.Second)

	timer.Stop()
	stopTimerCh <- true
	if err := sem.Release(); err != nil {
		log.Println(err)
		return
	}
	log.Println("release success")
}

func do2() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	log.Println(pong, err)

	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		uuid, err := uuid.NewV4()
		if err != nil {
			log.Println(err)
			return
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			do1(client, hex.EncodeToString(uuid.Bytes()))
		}()

		time.Sleep(11 * time.Second)
	}

	wg.Wait()
}

// 测试 zcard 压力
func do3() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	for i := 1; i <= 1000000; i++ {
		uuid, err := uuid.NewV4()
		if err != nil {
			log.Println(err)
			return
		}

		mem := &redis.Z{
			Score:  float64(time.Now().UnixNano()),
			Member: hex.EncodeToString(uuid.Bytes()),
		}
		_, err = client.ZAdd("zset-zcard-press", mem).Result()
		if err != nil {
			log.Println(err)
			return
		}
		//log.Println("zadd result:", n)

		if i%100000 == 0 {
			start := time.Now()
			zcard, err := client.ZCard("zset-zcard-press").Result()
			if err != nil {
				log.Println(err)
				return
			}
			log.Printf("zcard result: %d cost: %d ns\n", zcard, time.Now().Sub(start).Nanoseconds())
		}
	}
}

// 测试 zremrangebyscore 压力
func do4() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	for i := 1; i <= 1000000; i++ {
		uuid, err := uuid.NewV4()
		if err != nil {
			log.Println(err)
			return
		}

		mem := &redis.Z{
			Score:  float64(i),
			Member: hex.EncodeToString(uuid.Bytes()),
		}
		_, err = client.ZAdd("zset-zremrangebyscore-press", mem).Result()
		if err != nil {
			log.Println(err)
			return
		}
		//log.Println("zadd result:", n)
	}

	start := time.Now()
	zremrangebyscore, err := client.ZRemRangeByScore("zset-zremrangebyscore-press", "0", "100000").Result()
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("100000 zremrangebyscore result: %d cost: %d ms\n", zremrangebyscore, time.Now().Sub(start).Nanoseconds()/1000000)

	start = time.Now()
	zremrangebyscore, err = client.ZRemRangeByScore("zset-zremrangebyscore-press", "100001", "300000").Result()
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("200000 zremrangebyscore result: %d cost: %d ms\n", zremrangebyscore, time.Now().Sub(start).Nanoseconds()/1000000)

	start = time.Now()
	zremrangebyscore, err = client.ZRemRangeByScore("zset-zremrangebyscore-press", "300001", "600000").Result()
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("300000 zremrangebyscore result: %d cost: %d ms\n", zremrangebyscore, time.Now().Sub(start).Nanoseconds()/1000000)

	start = time.Now()
	zremrangebyscore, err = client.ZRemRangeByScore("zset-zremrangebyscore-press", "600001", "1000000").Result()
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("400000 zremrangebyscore result: %d cost: %d ms\n", zremrangebyscore, time.Now().Sub(start).Nanoseconds()/1000000)
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do2()
}
