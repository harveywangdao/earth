package main

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

/*
http://localhost:7849

Connection: keep-alive
Connection: close
*/
func send(client *http.Client) {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	//req.Header.Set("Connection", "keep-alive")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Println("http request fail, http code:", resp.StatusCode)
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("body:", string(data))
}

func c1() {
	client := &http.Client{}
	co := make(chan int, 10)
	for {
		co <- 1
		log.Println("并发数:", len(co))
		go func() {
			defer func() {
				<-co
			}()
			send(client)
		}()
	}
}

func c2() {
	co := make(chan int, 10)
	for {
		co <- 1
		//log.Println("并发数:", len(co))
		go func() {
			defer func() {
				<-co
			}()
			client := &http.Client{}
			send(client)
		}()
	}
}

//复用1条tcp链接
func c3() {
	client := &http.Client{
		Transport: &http.Transport{},
	}
	for i := 0; i < 10; i++ {
		send(client)
	}
}

//不复用1条tcp链接
func c4() {
	for i := 0; i < 10; i++ {
		client := &http.Client{
			Transport: &http.Transport{},
		}
		send(client)
	}
}

//复用10条tcp链接
func c5() {
	client := &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConnsPerHost: 10,
			MaxIdleConns:        10,
			IdleConnTimeout:     60 * time.Second,
		},
		Timeout: 600 * time.Second,
	}

	for i := 1; i <= 50; i++ {
		go func() {
			send(client)
		}()

		if i%10 == 0 {
			time.Sleep(30 * time.Second)
		}
	}
}

//无法复用10条tcp链接
func c6() {
	for i := 1; i <= 50; i++ {
		go func() {
			client := &http.Client{
				Transport: &http.Transport{
					DialContext: (&net.Dialer{
						Timeout:   30 * time.Second,
						KeepAlive: 30 * time.Second,
					}).DialContext,
					MaxIdleConnsPerHost: 10,
					MaxIdleConns:        10,
					IdleConnTimeout:     60 * time.Second,
				},
				Timeout: 600 * time.Second,
			}
			send(client)
		}()

		if i%10 == 0 {
			time.Sleep(30 * time.Second)
		}
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	c4()
	time.Sleep(30 * time.Second)
}
