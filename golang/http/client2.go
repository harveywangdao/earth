package main

import (
	"bytes"
	"fmt"
	//"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	//"os"
)

func main() {
	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)

	/*fileWriter1, _ := bodyWriter.CreateFormFile("files", "file1.txt")
	file1, _ := os.Open("file1.txt")
	defer file1.Close()
	io.Copy(fileWriter1, file1)

	fileWriter2, _ := bodyWriter.CreateFormFile("files", "file2.txt")
	file2, _ := os.Open("file2.txt")
	defer file2.Close()
	io.Copy(fileWriter2, file2)*/

	// other form data
	extraParams := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	for key, value := range extraParams {
		_ = bodyWriter.WriteField(key, value)
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	fmt.Println(bodyBuffer.String())

	resp, err := http.Post("http://localhost:5616/upload", contentType, bodyBuffer)
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
