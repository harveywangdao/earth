package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func do1() {
	r := gin.Default()

	r.GET("/hoho", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hoho",
		})
	})

	r.Use(func(c *gin.Context) {
		log.Println("111")
		c.Next()
	})

	r.GET("/haha", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "haha",
		})
	})

	table := r.Group("/api")
	{
		table.GET("/hehe", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "hehe",
			})
		})

		table.Use(func(c *gin.Context) {
			log.Println("222")
			c.Next()
		})

		table.GET("/sese", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "sese",
			})
		})
	}

	r.Run(":9990")
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	do1()
}
