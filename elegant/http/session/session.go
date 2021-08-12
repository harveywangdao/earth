package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/go-session/session"

	gsessions "github.com/gorilla/sessions"
)

func do1() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		store, err := session.Start(context.Background(), w, r)
		if err != nil {
			fmt.Fprint(w, err)
			log.Println(err)
			return
		}

		store.Set("foo", "bar")
		if err := store.Save(); err != nil {
			fmt.Fprint(w, err)
			log.Println(err)
			return
		}

		http.Redirect(w, r, "/foo", http.StatusFound)
	})

	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		store, err := session.Start(context.Background(), w, r)
		if err != nil {
			fmt.Fprint(w, err)
			log.Println(err)
			return
		}

		foo, ok := store.Get("foo")
		if ok {
			fmt.Fprintf(w, "foo: %s", foo)
			log.Printf("foo: %s", foo)
			return
		}
		log.Println("foo not exist")
		fmt.Fprint(w, "foo not exist\n")
	})

	log.Println("http server listen:", "8564")
	log.Fatal(http.ListenAndServe(":8564", mux))
}

func do2() {
	mux := http.NewServeMux()
	store := gsessions.NewCookieStore([]byte("asdasdasd"))

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session-name")
		log.Println("session:", session.Values)

		session.Values["foo"] = "bar"
		session.Values[42] = 44
		err := session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, "hello\n")
	})

	log.Println("http server listen:", "8564")
	log.Fatal(http.ListenAndServe(":8564", mux))
}

func do3() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/hello", func(c *gin.Context) {
		session := sessions.Default(c)

		if session.Get("hello") != "world" {
			session.Set("hello", "world")
			session.Save()
			log.Println("create session")
		}

		c.JSON(200, gin.H{"hello": session.Get("hello")})
	})
	r.Run(":8564")
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do3()
}
