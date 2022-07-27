package main

import (
	"encoding/json"
	"fmt"
	ipfs "github.com/ipfs/go-ipfs-api"
	"os"
	"strings"
	"time"
)

func main() {
	sh := ipfs.NewShell("localhost:5001")

	//id
	nodeInfo, err := sh.ID()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	jsonData, _ := json.Marshal(nodeInfo)

	fmt.Println(string(jsonData))

	//Add string
	cid, err := sh.Add(strings.NewReader("hello worldfddf!"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("added", cid)

	//Cat string
	data := make([]byte, 1024)
	for {
		time.Sleep(1 * time.Second)
		resp, err := sh.Cat(cid)
		if err != nil {
			fmt.Println(err)
			continue
		}

		n, err := resp.Read(data)
		resp.Close()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(cid, ":", string(data[:n]))
		break
	}

	//Add file
	filePath := "./aaa.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileHash, err := sh.Add(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("added file", fileHash)

	//Cat file
	for {
		time.Sleep(1 * time.Second)
		resp, err := sh.Cat(fileHash)
		if err != nil {
			fmt.Println(err)
			continue
		}

		n, err := resp.Read(data)
		resp.Close()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(fileHash, ":", string(data[:n]))
		break
	}

	//Add dir
	dir := "./testing"
	dirHash, err := sh.AddDir(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("added dir", dirHash)

	links, err := sh.List(dirHash)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	jsonData, _ = json.Marshal(links)

	fmt.Println(string(jsonData))

	for i := 0; i < len(links); i++ {
		if links[i].Type == 1 {

		} else if links[i].Type == 2 {
			for {
				time.Sleep(1 * time.Second)
				resp, err := sh.Cat(links[i].Hash)
				if err != nil {
					fmt.Println(err)
					continue
				}

				n, err := resp.Read(data)
				resp.Close()
				if err != nil {
					fmt.Println(err)
					continue
				}

				fmt.Println(links[i].Name, ":", string(data[:n]))
				break
			}
		} else {
			fmt.Println("file type error")
			os.Exit(1)
		}
	}
}
