package main

import (
	"io"
	"mime/multipart"

	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	piper, pipew := io.Pipe()
	defer piper.Close()

	bodyWriter := multipart.NewWriter(pipew)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer pipew.Close()

		bodyWriter.WriteField("key1", "value1")

		part, err := bodyWriter.CreateFormFile("files", "bigfile.tar")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer bodyWriter.Close()

		file, err := os.Open("/home/thomas/download/bigfile.tar")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		if _, err = io.Copy(part, file); err != nil {
			fmt.Println(err)
			return
		}
	}()

	fmt.Println("post start")
	resp, err := http.Post("http://localhost:5616/upload", bodyWriter.FormDataContentType(), piper)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("post end")

	fmt.Println("read body start")
	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("read body end")

	fmt.Println(string(ret))

	wg.Wait()
}
