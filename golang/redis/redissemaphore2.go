package main

import (
	"fmt"
	"github.com/go-redis/redis"
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

	stopTimerCh        chan bool
	refreshTimer       *time.Timer
	refreshTimerPeriod time.Duration // 刷新定时器的间隔时间

	mu       sync.Mutex
	unlocked bool
}

func RedisSemaphoreTryWait(cli *redis.Client, semKey, semMember string, maxSemCount int64, semTimeout time.Duration) (*RedisSemaphore, error) {
	// 清除过期的信号
	// 获取信号数量,如果小于最大信号量就增加一个信号量,否则返回失败
	script := redis.NewScript(`
    redis.call('zremrangebyscore', KEYS[1], '-inf', ARGV[2])
    if redis.call('zcard', KEYS[1]) < tonumber(ARGV[3]) then
      return redis.call('zadd', KEYS[1], ARGV[1], ARGV[4], ARGV[5])
    else
      return -1
    end`)

	r := &RedisSemaphore{
		Client:      cli,
		SemKey:      semKey,
		SemMember:   semMember,
		MaxSemCount: maxSemCount,
		SemTimeout:  semTimeout,
	}

	start := time.Now()
	defer func() {
		log.Printf("get redis semaphore cost %d ms", time.Now().Sub(start).Nanoseconds()/1000000)
	}()

	n, err := script.Run(r.Client, []string{r.SemKey}, "nx", time.Now().Add(-1*r.SemTimeout).UnixNano(), r.MaxSemCount, time.Now().UnixNano(), r.SemMember).Int64()
	if err != nil {
		return nil, err
	} else if n < 0 {
		return nil, fmt.Errorf("get sem fail")
	} else if n == 0 {
		return nil, fmt.Errorf("sem already existed")
	}

	r.refreshTimerPeriod = (r.SemTimeout * 8) / 10
	r.refreshTimer = time.NewTimer(r.refreshTimerPeriod)
	r.stopTimerCh = make(chan bool)
	go func() {
		for {
			select {
			case <-r.refreshTimer.C:
				if err := r.Refresh(); err != nil {
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

	n, err := r.Client.ZRem(r.SemKey, r.SemMember).Result()
	if err != nil {
		return err
	} else if n != 1 {
		return fmt.Errorf("sem: %s member: %s not existed", r.SemKey, r.SemMember)
	}
	return nil
}

func do2() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	sem, err := RedisSemaphoreTryWait(client, "redis-sem", "semMember", 10, 10*time.Second)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("get sem success")

	go func() {
		sem, err := RedisSemaphoreTryWait(client, "redis-sem", "semMember2", 10, 10*time.Second)
		if err != nil {
			log.Println(err)
			return
		}
		defer sem.Release()
		log.Println("get sem2 success")
	}()

	time.Sleep(30 * time.Second)

	if err := sem.Release(); err != nil {
		log.Println(err)
		return
	}
	log.Println("release success")
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do2()
}
