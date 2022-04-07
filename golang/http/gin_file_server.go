package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()

	router.POST("/upload/:taskid/:filename", func(c *gin.Context) {
		file, err := c.FormFile("files")
		if err != nil {
			c.String(http.StatusBadRequest, "请求失败")
			return
		}

		fmt.Println("taskid:", c.Param("taskid"))
		fmt.Println("filename:", c.Param("filename"))

		os.Mkdir("/home/thomas/test/files/"+c.Param("taskid"), os.ModePerm)

		if err := c.SaveUploadedFile(file, "/home/thomas/test/files/"+c.Param("taskid")+"/"+c.Param("filename")); err != nil {
			c.String(http.StatusBadRequest, "保存失败 Error:%s", err.Error())
			return
		}

		data, err := ioutil.ReadFile("/home/thomas/test/files/" + c.Param("taskid") + "/" + c.Param("filename"))
		if err != nil {
			c.String(http.StatusBadRequest, "读取文件失败")
			return
		}

		fileMd5 := md5.Sum(data)
		fileMd5Base64 := base64.StdEncoding.EncodeToString(fileMd5[:])
		fmt.Println(fileMd5, fileMd5Base64, hex.EncodeToString(fileMd5[:]))

		rfileMd5Base64 := c.Request.Header.Get("File-MD5")
		rfileMd5, err := base64.StdEncoding.DecodeString(rfileMd5Base64)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		fmt.Println(rfileMd5, rfileMd5Base64, hex.EncodeToString(rfileMd5))

		c.String(http.StatusOK, "上传文件成功")
	})

	router.Run(":7845")
}
