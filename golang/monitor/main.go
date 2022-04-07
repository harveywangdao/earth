package main

import (
	"log"
	"net/http"

	"github.com/arl/statsviz"
	"github.com/gin-gonic/gin"
)

func do1() {
	mux := http.NewServeMux()
	statsviz.Register(mux)
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong\n"))
	})
	server := &http.Server{Addr: ":8085", Handler: mux}
	log.Println("golang http server start")
	log.Fatal(server.ListenAndServe())
}

func do2() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/debug/statsviz/*filepath", func(c *gin.Context) {
		if c.Param("filepath") == "/ws" {
			statsviz.Ws(c.Writer, c.Request)
			return
		}
		statsviz.IndexAtRoot("/debug/statsviz").ServeHTTP(c.Writer, c.Request)
	})

	log.Println("golang gin server start")
	r.Run(":8086")
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	go do1()
	go do2()
	select {}
}
