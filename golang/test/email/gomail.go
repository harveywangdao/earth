package main

import (
	"fmt"
	gomail "gopkg.in/gomail.v2"
)

func main() {
	fmt.Println("Send email start")

	m := gomail.NewMessage()
	m.SetAddressHeader("From", "3227855932@qq.com", "Thomas")
	m.SetAddressHeader("To", "thomasdarwin@163.com", "Darwin")
	m.SetAddressHeader("Cc", "1272777053@qq.com", "Harvey")
	m.SetAddressHeader("Bcc", "1807026130@qq.com", "Christophe")
	m.SetHeader("Subject", "Golang Conference Invitation")
	m.SetBody("text/html", "Hello New Email.<br>Hello New Email.")
	m.Attach("/home/thomas/study/golang/webserver/beego/thirdparty/src/email/main.go")

	d := gomail.NewDialer("smtp.qq.com", 465, "3227855932@qq.com", "gssgyijwteigdahb")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	//TLS/SSL Authorization code
	//smtp.qq.com:465
	//1272777053@qq.com vsdunuxkyesyheji
	//3227855932@qq.com gssgyijwteigdahb
	//thomasdarwin@163.com 180498dao
	fmt.Println("Send email finish")
}
