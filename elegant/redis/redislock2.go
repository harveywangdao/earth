package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/bsm/redislock"
	"github.com/go-redis/redis"
)

// 可以自动刷新的redis分布式锁
type RedisDistributedLock struct {
	cli         *redis.Client
	key         string
	keyTTL      time.Duration // redis key过期时间
	retryPeriod time.Duration // 锁失败后的重试间隔时间

	locker *redislock.Lock

	stopTimerCh        chan bool
	refreshTimer       *time.Timer
	refreshTimerPeriod time.Duration // 刷新定时器的间隔时间

	mu       sync.Mutex
	unlocked bool
}

func RedisDistLock(cli *redis.Client, key string, ttl time.Duration, isBlock bool) (*RedisDistributedLock, error) {
	r := &RedisDistributedLock{
		cli: cli,
		key: fmt.Sprintf("redis-dist-%s", key),

		refreshTimerPeriod: (ttl * 8) / 10,
		keyTTL:             ttl,
	}

	// 阻塞代表第一次获取key失败会重试多次,直到超时
	if isBlock {
		r.retryPeriod = ttl / 10
	}

	var err error
	option := &redislock.Options{
		RetryStrategy: redislock.LinearBackoff(r.retryPeriod),
	}
	r.locker, err = redislock.Obtain(r.cli, r.key, r.keyTTL, option)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	r.refreshTimer = time.NewTimer(r.refreshTimerPeriod)
	r.stopTimerCh = make(chan bool)
	go func() {
		for {
			select {
			case <-r.refreshTimer.C:
				if err := r.locker.Refresh(r.keyTTL, nil); err != nil {
					log.Println(err)
					continue
				}
				log.Println("refresh success")
				r.refreshTimer.Reset(r.refreshTimerPeriod)
			case <-r.stopTimerCh:
				log.Println("stopTimerCh")
				return
			}
		}
	}()

	return r, nil
}

func (r *RedisDistributedLock) Unlock() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// 防止unlock多次
	if r.unlocked {
		return fmt.Errorf("already unlocked")
	}

	r.unlocked = true
	if r.refreshTimer != nil {
		r.refreshTimer.Stop()
	}
	if r.stopTimerCh != nil {
		r.stopTimerCh <- true
	}

	if err := r.locker.Release(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func do1() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	locker, err := RedisDistLock(client, "appid01", 10*time.Second, false)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("get locker success")

	time.Sleep(5 * time.Second)
	go func() {
		locker2, err := RedisDistLock(client, "appid01", 10*time.Second, false)
		if err != nil {
			log.Println(err)
			return
		}
		defer locker2.Unlock()
		log.Println("get locker2 success")
	}()

	time.Sleep(25 * time.Second)

	if err := locker.Unlock(); err != nil {
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

	locker, err := RedisDistLock(client, "appid01", 10*time.Second, true)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("get locker success")

	time.Sleep(5 * time.Second)
	go func() {
		locker2, err := RedisDistLock(client, "appid01", 10*time.Second, true)
		if err != nil {
			log.Println(err)
			return
		}
		defer locker2.Unlock()
		log.Println("get locker2 success")
	}()

	time.Sleep(25 * time.Second)

	if err := locker.Unlock(); err != nil {
		log.Println(err)
		return
	}
	log.Println("release success")
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do2()
	time.Sleep(1 * time.Hour)
}
