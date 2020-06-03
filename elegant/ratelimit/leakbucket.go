package main

import (
	"fmt"
	"time"
)

// 桶
type LeakBucket struct {
	capacity       int           // 容量，固定时间语序访问次数
	interval       time.Duration // 允许访问的时间间隔
	dropsNum       int           // 固定时间访问了多少次
	lastAccessTime time.Time     // 最近一次的访问时间
}

func (b *LeakBucket) accessControl() bool {
	now := time.Now()
	pastTime := now.Sub(b.lastAccessTime)

	// 过去这段时间pastTime可以允许多少访问
	leaks := int(float64(pastTime) / float64(b.interval))
	if leaks > 0 { // 说明这段时间可以有leaks可以访问，但没有用户访问
		// 所以放宽访问，下一段访问限制减少一定leaks次限制，通过这种机制达到平滑控制流量
		fmt.Println("leaks:", leaks)

		if b.dropsNum <= leaks {
			b.dropsNum = 0
		} else {
			b.dropsNum -= leaks
		}
		// 更新访问时间
		b.lastAccessTime = now
	}

	if b.dropsNum < b.capacity { // 在允许访问之内
		b.dropsNum++
		return true
	} else {
		return false
	}
}

func main() {
	bucket := &LeakBucket{
		capacity: 1,
		interval: time.Second,
	}

	count := 0
	start := time.Now()
	for i := 0; i < 10000; i++ {
		allowed := bucket.accessControl()
		if allowed {
			fmt.Println("i", i)
			fmt.Println("allowed", allowed)
			count++
		}

		time.Sleep(time.Millisecond)
	}

	fmt.Println("count:", count)
	fmt.Println(time.Now().Sub(start).Seconds())
}
