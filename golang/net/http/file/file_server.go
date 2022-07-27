package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func uploadHandler1(w http.ResponseWriter, r *http.Request) {
	reader, err := r.MultipartReader()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}

		fmt.Printf("FileName=[%s], FormName=[%s]\n", part.FileName(), part.FormName())

		if part.FileName() == "" {
			data, _ := ioutil.ReadAll(part)
			fmt.Printf("FormData=[%s]\n", string(data))
		} else {
			dst, _ := os.Create("./" + part.FileName() + ".upload")
			defer dst.Close()
			io.Copy(dst, part)
		}
	}
}

func uploadHandler2(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("files")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("header.Filename:", header.Filename)

	fmt.Println(r.FormValue("key1"))
	fmt.Println(r.FormValue("key2"))
	fmt.Println(r.FormValue("key3"))
	fmt.Println(r.FormValue("key4"))

	fmt.Println(r.PostFormValue("key1"))
	fmt.Println(r.PostFormValue("key2"))
	fmt.Println(r.PostFormValue("key3"))
	fmt.Println(r.PostFormValue("key4"))

	dst, _ := os.Create("./" + header.Filename + ".upload")
	defer dst.Close()
	io.Copy(dst, file)
}

func uploadHandler3(w http.ResponseWriter, r *http.Request) {
	//r.ParseMultipartForm(maxUploadSize)
	//defer r.MultipartForm.RemoveAll()

	fmt.Println("get file start")
	file, header, err := r.FormFile("files")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Println("get file end")

	fmt.Println("get value start")
	fmt.Println(r.PostFormValue("key1"))
	fmt.Println("get value end")

	dst, err := os.Create("/home/thomas/download/" + header.Filename + ".upload")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dst.Close()

	fmt.Println("copy start")
	_, err = io.Copy(dst, file)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("copy end")
}

func main() {
	http.HandleFunc("/upload", uploadHandler3)
	http.ListenAndServe(":5616", nil)
}
