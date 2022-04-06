package main

import (
	"fmt"
	"sync"
)

type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		set.RLock()

		for _, elem := range set.s {
			ch <- elem
		}

		close(ch)
		set.RUnlock()
	}()

	return ch
}

func main() {
	th := threadSafeSet{
		s: []interface{}{"1", "2"},
	}

	v := <-th.Iter()

	fmt.Printf("%s%v\n", "ch", v)
}
