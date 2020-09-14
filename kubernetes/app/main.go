package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func metrics() {
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":7878", nil))
}

func main() {
	var dir, dir2 string
	flag.StringVar(&dir, "dir", "/data/pear", "dir")
	flag.StringVar(&dir2, "dir2", "/data/pear_pvc", "dir2")
	flag.Parse()

	go metrics()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		err := os.MkdirAll(dir, 0777)
		if err != nil {
			c.JSON(200, gin.H{
				"message": err.Error(),
			})
			return
		}

		f, err := os.OpenFile(filepath.Join(dir, "data.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			c.JSON(200, gin.H{
				"message": err.Error(),
			})
			return
		}
		defer f.Close()

		_, err = f.WriteString(time.Now().Format("Mon Jan 2 15:04:05 -0700 MST 2006") + "\n")
		if err != nil {
			c.JSON(200, gin.H{
				"message": err.Error(),
			})
			return
		}

		err = os.MkdirAll(dir2, 0777)
		if err != nil {
			c.JSON(200, gin.H{
				"message": err.Error(),
			})
			return
		}

		f2, err := os.OpenFile(filepath.Join(dir2, "data.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			c.JSON(200, gin.H{
				"message": err.Error(),
			})
			return
		}
		defer f2.Close()

		_, err = f2.WriteString(time.Now().Format("Mon Jan 2 15:04:05 -0700 MST 2006") + "\n")
		if err != nil {
			c.JSON(200, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "today is yesterday",
		})
	})
	r.Run(":6666") // listen and serve on 0.0.0.0:8080
}
