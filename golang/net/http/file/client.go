package main

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)

	fileWriter, _ := bodyWriter.CreateFormFile("files", "file1.txt")
	file, _ := os.Open("file1.txt")
	defer file.Close()
	io.Copy(fileWriter, file)

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	req, err := http.NewRequest(http.MethodPost, "http://localhost:7845/upload/saaa/1.png", bodyBuffer)
	if err != nil {
		log.Fatal(err)
		return
	}
	req.Header.Set("Content-Type", contentType)

	file.Seek(0, os.SEEK_SET)
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(string(data))

	md5Data := md5.Sum(data)
	req.Header.Set("file-md5", base64.StdEncoding.EncodeToString(md5Data[:]))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println(resp.Status)
	log.Println(string(resp_body))
}
