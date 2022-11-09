package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

//r.FormFile(key)     // r.MultipartForm
//r.MultipartReader() // r.MultipartForm
//r.FormValue(key)     // r.Form
//r.PostFormValue(key) // r.PostForm
// 多文件+其它值
func do1(w http.ResponseWriter, r *http.Request) {
	multipartReader, err := r.MultipartReader()
	if err != nil {
		log.Fatal(err)
	}
	defer r.MultipartForm.RemoveAll()

	for {
		part, err := multipartReader.NextPart()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		//part.Close()
		log.Println(part.FileName())
		log.Println(part.FormName())
		data, err := ioutil.ReadAll(part)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(data))
	}
}

func do2(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("fag1") // 同一个key可以传多个文件,但是FormFile只能读取第一个文件
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	defer r.MultipartForm.RemoveAll()

	log.Println(header.Filename)
	log.Println(header.Header)
	log.Println(header.Size)

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(data))
}

func do3(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("people")
	log.Println(v)
	v2 := r.PostFormValue("people")
	log.Println(v2)
	r.MultipartForm.RemoveAll()
}

func do4(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/octet-stream")
	//w.Header().Set("Content-Disposition", "attachment")
	//w.Header().Set("Content-Disposition", "attachment;filename=1.docx")
	//w.Header().Set("Content-Disposition", "attachment;filename=我是.docx")
	//w.Header().Set("Content-Disposition", "attachment;filename="+url.QueryEscape("我是.docx"))
	w.Header().Set("Content-Disposition", "attachment;filename="+url.PathEscape("我是.docx"))
	w.WriteHeader(http.StatusOK)

	data := []byte("125dfg")
	_, err := io.Copy(w, bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	http.HandleFunc("/do1", do1)
	http.HandleFunc("/do2", do2)
	http.HandleFunc("/do3", do3)
	http.HandleFunc("/do4", do4)
	http.ListenAndServe(":5616", nil)
}
