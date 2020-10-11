package main

import (
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	db, err := leveldb.OpenFile("xiaoming", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	data, err := db.Get([]byte("key02"), nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("data1:", string(data))

	err = db.Put([]byte("key02"), []byte("value02"), nil)
	if err != nil {
		log.Fatal(err)
	}

	data, err = db.Get([]byte("key02"), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("data2:", string(data))

	err = db.Delete([]byte("key02"), nil)
	if err != nil {
		log.Fatal(err)
	}

	data, err = db.Get([]byte("key02"), nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("data3:", string(data))
}
