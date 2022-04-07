package main

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/oklog/run"
	//"golang.org/x/sync/errgroup"
)

func myProcess(ctx context.Context) error {
	log.Println("start")
	time.Sleep(2 * time.Second)
	log.Println("start")

	return nil
}

func do1() {
	var g run.Group
	ctx, cancel := context.WithCancel(context.Background())
	g.Add(func() error {
		return myProcess(ctx)
	}, func(err error) {
		log.Println(err)
		cancel()
	})

	g.Run()
}

func do2() {
	var g run.Group
	{
		cancel := make(chan struct{})
		g.Add(func() error {
			select {
			case <-time.After(time.Second):
				log.Printf("The first actor had its time elapsed\n")
				return nil
			case <-cancel:
				log.Printf("The first actor was canceled\n")
				return nil
			}
		}, func(err error) {
			log.Printf("The first actor was interrupted with: %v\n", err)
			close(cancel)
		})
	}
	{
		g.Add(func() error {
			log.Printf("The second actor is returning immediately\n")
			return errors.New("immediate teardown")
		}, func(err error) {
			// Note that this interrupt function is called, even though the
			// corresponding execute function has already returned.
			log.Printf("The second actor was interrupted with: %v\n", err)
		})
	}
	log.Printf("The group was terminated with: %v\n", g.Run())
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do2()
}
