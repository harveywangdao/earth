package controllers

import (
	"github.com/astaxie/beego"
	"zeus/models"
)

type ReplyController struct {
	beego.Controller
}

func (c *ReplyController) Get() {

}

func (c *ReplyController) Post() {

}

func (c *ReplyController) Add() {
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}

	tid := c.Input().Get("tid")
	nickname := c.Input().Get("nickname")
	replycontent := c.Input().Get("replycontent")

	err := models.AddReply(tid, nickname, replycontent)
	if err != nil {
		beego.Error(err)
	}

	c.Redirect("/topic/view/"+tid, 302)
}

func (c *ReplyController) Modify() {

}

func (c *ReplyController) Delete() {
	beego.Trace("Delete")
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}

	id := c.Input().Get("id")
	tid := c.Input().Get("tid")

	err := models.DeleteOneReply(id)
	if err != nil {
		beego.Error(err)
	}

	c.Redirect("/topic/view/"+tid, 302)
}
