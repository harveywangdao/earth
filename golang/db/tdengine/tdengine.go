package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/taosdata/driver-go/v3/taosWS"
)

func test01() {
	var taosUri = "root:taosdata@ws(localhost:6041)/"
	taos, err := sql.Open("taosWS", taosUri)
	if err != nil {
		log.Println("failed to connect TDengine, err:", err)
		return
	}
	defer taos.Close()
	taos.Exec("create database if not exists test")
	taos.Exec("use test")
	taos.Exec("create table if not exists tb1 (ts timestamp, a int)")
	_, err = taos.Exec("insert into tb1 values(now, 0)(now+1s,1)(now+2s,2)(now+3s,3)")
	if err != nil {
		log.Println("failed to insert, err:", err)
		return
	}
	rows, err := taos.Query("select * from tb1")
	if err != nil {
		log.Println("failed to select from table, err:", err)
		return
	}
	defer rows.Close()
	log.Println("sdfs")
	for rows.Next() {
		var r struct {
			ts time.Time
			a  int
		}
		err := rows.Scan(&r.ts, &r.a)
		if err != nil {
			log.Println("scan error:", err)
			return
		}
		log.Println(r.ts, r.a)
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	test01()
}
