package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func do1() {
	buf := bytes.NewBufferString("hello, I am http client")
	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8564/api", buf)
	if err != nil {
		log.Fatal(err)
	}
	cli := &http.Client{}

	resp, err := cli.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	log.Println(resp.Header)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("http client recv msg: %s", data)

	if resp.StatusCode != http.StatusOK {
		log.Fatal("http code:", resp.Status)
	}
}

func do2() {
	buf := bytes.NewBufferString("hello, I am http client")
	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8564/api", buf)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Connection", "keep-alive")
	//req.Header.Set("Connection", "close")

	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	//defer resp.Body.Close()

	log.Println(resp.Header)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("http client recv msg: %s", data)

	if resp.StatusCode != http.StatusOK {
		log.Fatal("http code:", resp.Status)
	}

	resp.Body.Close()
	time.Sleep(time.Hour)
}

func do3() {
	cli := &http.Client{}

	for {
		buf := bytes.NewBufferString("hello, I am http client")
		req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8564/api", buf)
		if err != nil {
			log.Fatal(err)
		}

		//req.Header.Set("Connection", "keep-alive")
		req.Header.Set("Connection", "close")

		resp, err := cli.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(resp.Header)

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("http client recv msg: %s", data)

		if resp.StatusCode != http.StatusOK {
			log.Fatal("http code:", resp.Status)
		}
		resp.Body.Close()

		time.Sleep(time.Second)
	}
}

func do4() {
	bufr, bufw := io.Pipe()

	go func() {
		bufw.Write([]byte("orange1"))
		time.Sleep(time.Second)
		bufw.Write([]byte("orange2"))
		time.Sleep(time.Second)
		bufw.Write([]byte("orange3"))
		time.Sleep(time.Second)
		bufw.Write([]byte("orange4"))
		time.Sleep(time.Second)
		bufw.Write([]byte("orange5"))
		bufw.Close()
	}()

	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8564/chunk", bufr)
	if err != nil {
		log.Fatal(err)
	}
	//req.Header.Set("Transfer-Encoding", "chunked")

	log.Println("send http")

	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	log.Println(resp.Header)

	buf := make([]byte, 128)
	for {
		n, err := resp.Body.Read(buf)
		if err != nil {
			if err == io.EOF {
				if n > 0 {
					log.Printf("http client recv msg: %s", buf[:n])
				}
				break
			}
			log.Fatal(err)
		}
		log.Printf("http client recv msg: %s", buf[:n])
	}

	/*	data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("http client recv msg: %s", data)*/

	if resp.StatusCode != http.StatusOK {
		log.Fatal("http code:", resp.Status)
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do4()
}
