package main

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", Cookie)
	http.HandleFunc("/2", Cookie2)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Cookie(w http.ResponseWriter, r *http.Request) {
	log.Print("111")
	ck := &http.Cookie{
		Name:  "MyCookie",
		Value: "hhell osss",
		Path:  "/",
		//Domain: "localhost",
		MaxAge: 120,
	}

	http.SetCookie(w, ck)

	ck2, err := r.Cookie("MyCookie")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	log.Print("222")
	io.WriteString(w, ck2.Value)
}

func Cookie2(w http.ResponseWriter, r *http.Request) {
	log.Print("111")
	ck := &http.Cookie{
		Name:  "MyCookie",
		Value: "hhelld  fffo",
		Path:  "/",
		//Domain: "localhost",
		MaxAge: 120,
	}

	//http.SetCookie(w, ck)
	w.Header().Set("Set-Cookie", strings.Replace(ck.String(), " ", "xx", -1))

	ck2, err := r.Cookie("MyCookie")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	log.Print("222")
	io.WriteString(w, ck2.Value)
}
