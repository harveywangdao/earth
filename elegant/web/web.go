package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

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

func main() {
	go ginweb()
	go beegoweb()
	//go irisweb()
	go echoweb()
	go fiberweb()
	go martiniweb()
	go chiweb()
	select {}
}
