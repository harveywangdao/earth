package main

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func m(c *gin.Context) {
	fmt.Println("Middleware start")

	fmt.Println(c.Request.Host)
	fmt.Println(c.Request.Header.Get("host"))
	fmt.Println(c.Request.URL)
	fmt.Println(c.Request.Referer())

	//c.Abort()
	//c.AbortWithError(http.StatusBadRequest, errors.New("alscfasca"))
	//c.AbortWithStatus(http.StatusBadRequest)
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": "bjabncjkanckmzn c",
	})
	return
	c.Next()
	fmt.Println("Middleware end")
}

func m1(c *gin.Context) {
	fmt.Println("Middleware1 start")
	c.Next()
	fmt.Println("Middleware1 end")
}

func m2(c *gin.Context) {
	fmt.Println("Middleware2 start")
	c.Next()
	fmt.Println("Middleware2 end")
}

func m3(c *gin.Context) {
	fmt.Println("Middleware3")
}

func m4(c *gin.Context) {
	fmt.Println("Middleware4")
}

func main() {
	r := gin.Default()

	r.Use(m)
	r.Use(m1)
	r.Use(m2)
	r.Use(m3)
	r.Use(m4)

	r.GET("/ping", func(c *gin.Context) {
		fmt.Println("handle...")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
