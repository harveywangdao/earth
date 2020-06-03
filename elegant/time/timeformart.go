package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().Unix()
	fmt.Println(t)

	fmt.Println(time.Unix(t, 0).Format("2006-01-02 15:04:05"))

	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.String())
	fmt.Println(now.UTC())
	fmt.Println(now.UTC().String())
	fmt.Println(now.Format("2006-01-02 15:04:05"))

	fmt.Println(now.AddDate(-1, 0, 0))

	t1, _ := time.Parse("2006-01-02 15:04:05 ", "2020-05-15 20:00:00")
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05 ", "2020-05-15 20:00:00", time.Local)
	fmt.Println("t1:", t1)
	fmt.Println("t2:", t2)
}
