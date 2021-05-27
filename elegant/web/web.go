package main

import (
	"crypto/tls"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	beego "github.com/beego/beego/server/web"
	bc "github.com/beego/beego/server/web/context"
	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-martini/martini"
	"github.com/gofiber/fiber"
	"github.com/kataras/iris"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func ginweb() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "gin web",
		})
	})
	r.Run(":9990")
}

func beegoweb() {
	beego.Get("/", func(ctx *bc.Context) {
		ctx.Output.JSON(map[string]string{"message": "beego web"}, false, false)
	})
	beego.Run(":9991")
}

func irisweb() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.JSON(map[string]string{"message": "iris web"})
	})
	app.Listen(":9992")
}

func echoweb() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "echo web"})
	})
	e.Start(":9993")
}

func fiberweb() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"message": "fiber web"})
	})
	app.Listen(":9994")
}

func martiniweb() {
	os.Setenv("PORT", "9995")
	m := martini.Classic()
	m.Get("/", func() string {
		resp := map[string]string{"message": "martini web"}
		data, _ := json.Marshal(resp)
		return string(data)
	})
	m.Run()
}

func chiweb() {
	r := chi.NewRouter()

	r.Use(chimiddleware.RequestID)
	r.Use(chimiddleware.RealIP)
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)
	r.Use(chimiddleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]string{"message": "chi web"}
		data, _ := json.Marshal(resp)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write(data)
	})
	http.ListenAndServe(":9996", r)
}

func https2web() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("https2 web\n"))
	})

	var srv http.Server
	srv.Addr = ":9997"
	srv.Handler = mux

	log.Fatal(srv.ListenAndServeTLS("../http/ca/server.crt", "../http/ca/server.key"))
}

func http2web() {
	go http2web_client()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("http2 web"))
	})
	s := &http.Server{
		Addr:    ":9998",
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}
	log.Fatal(s.ListenAndServe())
}

func http2web_client() {
	time.Sleep(time.Second * 5)
	client := &http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		},
	}
	resp, err := client.Get("http://localhost:9998")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	log.Println(resp.Status)
	log.Println(resp.Proto)
}

func http2web2() {
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("192.168.126.128"),
		Cache:      autocert.DirCache("certs"),
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("http2 web"))
	})

	srv := &http.Server{
		Addr: ":9998",
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}

	go http.ListenAndServe(":80", certManager.HTTPHandler(nil))

	log.Fatal(srv.ListenAndServeTLS("", ""))
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	go ginweb()
	go beegoweb()
	go irisweb()
	go echoweb()
	go fiberweb()
	go martiniweb()
	go chiweb()
	go https2web()
	go http2web()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	<-ch
}
