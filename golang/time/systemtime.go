package main

import (
	"log"
	"syscall"
	"time"
)

func do1() {
	t := time.Now()
	t = t.AddDate(1, 1, 1)
	tv := syscall.Timeval{
		Sec: t.Unix(),
	}
	if err := syscall.Settimeofday(&tv); err != nil {
		log.Printf("设置时间错误: %v", err)
	} else {
		log.Println("设置时间成功", t)
	}

	log.Println(time.Now())
	time.Sleep(time.Second * 20)
	log.Println(time.Now())
}

func main() {
	do1()
}
