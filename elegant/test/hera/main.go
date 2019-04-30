package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

/*func main() {
	http.HandleFunc("/", sayHello)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello beego web!")
}*/

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/hello", sayHello)

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(http.Dir(pwd))

	han2 := http.FileServer(http.Dir(pwd))

	//han1 := http.StripPrefix("/dddd/", han2)

	mux.Handle("/dddd/", han2)

	err = http.ListenAndServe(":8090", mux)

	if err != nil {
		log.Fatal(err)
	}
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "url:"+r.URL.String())
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	log.Fatal(r.URL.Path)
	io.WriteString(w, "Hello beego web!")
}

/*var muxtrouer map[string]func(http.ResponseWriter, *http.Request)

func main() {
	server := http.Server{
		Addr:        ":8090",
		Handler:     &myHandler{},
		ReadTimeout: 5 * time.Second,
	}

	muxtrouer = make(map[string]func(http.ResponseWriter, *http.Request))
	muxtrouer["/hello"] = sayHello
	muxtrouer["/bye"] = sayBye

	fmt.Println("step1")

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("step2")
	if h, ok := muxtrouer[r.URL.String()]; ok {
		fmt.Println("step4")
		h(w, r)
		return
	}
	fmt.Println("step4")
	io.WriteString(w, "url:"+r.URL.String())
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sayHello")
	io.WriteString(w, "Hello beego web!")
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sayBye")
	io.WriteString(w, "Bye beego web!")
}
*/
