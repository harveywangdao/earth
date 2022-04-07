package main

import (
	"fmt"
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
	"net/textproto"
)

func main() {
	fmt.Println("Send email start")
	e := &email.Email{
		To:      []string{"thomasdarwin@163.com"},
		From:    "Jordan Wright <3227855932@qq.com>",
		Subject: "Awesome Subject",
		Text:    []byte("Text Body is, of course, supported!"),
		HTML:    []byte("<h1>Fancy HTML is supported, too!</h1>"),
		Headers: textproto.MIMEHeader{},
	}

	err := e.Send("smtp.qq.com:465", smtp.PlainAuth("", "3227855932@qq.com", "gssgyijwteigdahb", "smtp.qq.com"))
	if err != nil {
		log.Fatal(err)
	}
	//1272777053@qq.com vsdunuxkyesyheji
	//3227855932@qq.com gssgyijwteigdahb
	//thomasdarwin@163.com 180498dao
	fmt.Println("Send email finish")
}
