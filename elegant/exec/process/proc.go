package main

import (
	"log"
	"os"
	"time"
)

func main() {
	process, err := os.StartProcess("./bottle", []string{}, &os.ProcAttr{})
	if err != nil {
		log.Fatal(err)
	}

	//log.Println("pid", process.Pid)
	//time.Sleep(time.Second * 15)
	log.Println("kill:", process.Kill())
	log.Println("kill:", process.Kill())
	//log.Println("release:", process.Release())

	log.Println("wait...")
	state, err := process.Wait()
	log.Println("wait result:", state.String(), err)

	time.Sleep(time.Minute)
}
