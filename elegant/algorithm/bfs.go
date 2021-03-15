package main

import (
	"container/list"
	"fmt"
)

func openLock1(deadends []string, target string) int {
	ts := []byte(target)
	steps := 0
	for i := 0; i < 4; i++ {
		n := int(ts[i] - '0')
		if n > 5 {
			n = 10 - n
		}
		steps += n
	}
	return steps
}

func openLock(deadends []string, target string) int {
	dm := make(map[string]int)
	for i := 0; i < len(deadends); i++ {
		dm[deadends[i]] = 1
	}
	if dm["0000"] == 1 {
		return -1
	}
	if target == "0000" {
		return 0
	}

	old := make(map[string]int)

	queue := list.New()
	old["0000"] = 1
	queue.PushFront("0000")

	step := 0
	for queue.Len() > 0 {
		for sz := queue.Len(); sz > 0; sz-- {
			e := queue.Back()
			queue.Remove(e)
			cur := []byte(e.Value.(string))

			for i := 0; i < 4; i++ {
				cur[i] = upone(cur[i])
				str := string(cur)
				if str == target {
					return step + 1
				}
				if dm[str] == 0 && old[str] == 0 {
					old[str] = 1
					queue.PushFront(str)
				}
				cur[i] = downone(cur[i])

				cur[i] = downone(cur[i])
				str = string(cur)
				if str == target {
					return step + 1
				}
				if dm[str] == 0 && old[str] == 0 {
					old[str] = 1
					queue.PushFront(str)
				}
				cur[i] = upone(cur[i])
			}
		}
		step++
	}
	return -1
}

func upone(c byte) byte {
	c++
	if c > '9' {
		return '0'
	}
	return c
}

func downone(c byte) byte {
	c--
	if c < '0' {
		return '9'
	}
	return c
}

func main() {
	fmt.Println(openLock(nil, "0202"))
}
