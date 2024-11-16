package main

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

func do1() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   "foo.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     1,    //days
		Compress:   true, // disabled by default
	})

	log.Println("lumberjack test")
}

func do2() {
	l := &lumberjack.Logger{
		Filename:   "/home/thomas/golang/src/test/cache/foo.log",
		MaxBackups: 2,
		//Compress: true,
	}
	log.SetOutput(l)

	for i := 0; i < 10; i++ {
		log.Println("lumberjack rotate", i)
		time.Sleep(time.Second)
		l.Rotate()
		fmt.Println("rotate", i)
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do2()
}
