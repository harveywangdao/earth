package main

import (
	"log"
	"net/http"
	"net/url"
)

func do1(s string) {
	cli := &http.Client{}

	url := "http://localhost:9990/" + s
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	resp, err := cli.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("http get fail, url: %s, code: %s", req.URL.String(), resp.Status)
		return
	}

	log.Printf("http get %s ok", req.URL.String())
}

func main() {
	s := "ss+ss"
	do1(url.PathEscape(s))
	do1(url.QueryEscape(s))

	s = "ss ss"
	do1(url.PathEscape(s))
	do1(url.QueryEscape(s))
}
