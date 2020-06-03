package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func removeLog(dir string, days int) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	for _, file := range files {
		if file.IsDir() {
			removeLog(filepath.Join(dir, file.Name()), days)
		} else {
			if file.ModTime().Add(time.Duration(days)*time.Hour*24).Unix() < time.Now().Unix() {
				fmt.Println(filepath.Join(dir, file.Name()), file.ModTime())
				os.Remove(filepath.Join(dir, file.Name()))
			}
		}
	}
}

func main() {
	//touch -d "2020-04-09 08:00:00" red/a.txt
	keepDay := "7"

	days, err := strconv.Atoi(keepDay)
	if err != nil {
		return
	}

	if days <= 0 {
		return
	}

	removeLog("/home/thomas/golang/src/test/red", days)
}
