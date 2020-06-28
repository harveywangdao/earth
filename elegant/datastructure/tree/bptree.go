package main

import (
	"fmt"
	"log"

	"github.com/timtadh/fs2/bptree"
	"github.com/timtadh/fs2/fmap"
	"github.com/timtadh/fs2/mmlist"
)

func do1() {
	bf, err := fmap.CreateBlockFile("haha.db")
	if err != nil {
		log.Fatal(err)
	}
	defer bf.Close()
	bpt, err := bptree.New(bf, 8, -1)
	if err != nil {
		log.Fatal(err)
	}
	_ = bpt
}

func do2() {
	bf, err := fmap.OpenBlockFile("haha.db")
	if err != nil {
		log.Fatal(err)
	}
	defer bf.Close()
	bpt, err := bptree.Open(bf)
	if err != nil {
		log.Fatal(err)
	}

	if err := bpt.Add([]byte("xiaoming"), []byte("xiaohong")); err != nil {
		log.Fatal(err)
	}

	kvi, err := bpt.Find([]byte("xiaoming"))
	if err != nil {
		log.Fatal(err)
	}
	for key, value, err, kvi := kvi(); kvi != nil; key, value, err, kvi = kvi() {
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(key), string(value))
	}
}

type Cup struct {
}

func (c *Cup) Write(p []byte) (n int, err error) {
	fmt.Println("aaaaa")
	return
}

func do3() {
	log.SetOutput(&Cup{})

	bf, err := fmap.OpenBlockFile("haha.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer bf.Close()
	bpt, err := bptree.Open(bf)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := bpt.Add([]byte("xiaoming"), []byte("xiaohong")); err != nil {
		log.Fatal(err)
	}

	key := []byte("xiaoming")

	err = bpt.DoFind(key, func(key, value []byte) error {
		fmt.Println("get key:", string(key), string(value))
		return nil
	})
	if err != nil {
		fmt.Println("DoFind fail:", err)
		return
	}

	err = bpt.Remove(key, func(value []byte) bool {
		fmt.Println("remove key:", string(key), string(value))
		return true
	})
	if err != nil {
		fmt.Println("Remove fail:", err)
		return
	}

	err = bpt.DoFind(key, func(key, value []byte) error {
		fmt.Println(string(key), string(value))
		return nil
	})
	if err != nil {
		fmt.Println("DoFind fail:", err)
		return
	}
}

func do4() {
	file, err := fmap.CreateBlockFile("bag.db")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	list, err := mmlist.New(file)
	if err != nil {
		log.Fatal(err)
		return
	}

	idx, err := list.Append([]byte("hello"))
	if err != nil {
		log.Println(err)
		return
	}

	value, err := list.Get(idx)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("get:", idx, string(value))

	if err := list.Set(idx, []byte("bye!")); err != nil {
		log.Println(err)
		return
	}

	value, err = list.Get(idx)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("get:", idx, string(value))

	value, err = list.Pop()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("pop:", string(value))

	value, err = list.Pop()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("pop:", string(value))
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do4()
}
