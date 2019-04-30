package controllers

import (
	"github.com/astaxie/beego"
	"path"
	"strings"
	"zeus/models"
)

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get() {
	c.Data["IsTopic"] = true
	c.TplName = "topic.html"
	c.Data["IsLogined"] = checkAccount(c.Ctx)

	var err error
	c.Data["Topics"], err = models.GetAllTopics("", "", false)
	if err != nil {
		beego.Error(err)
	}
}

func (c *TopicController) Post() {
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}

	title := c.Input().Get("title")
	content := c.Input().Get("content")
	topicCategory := c.Input().Get("topiccategory")
	label := c.Input().Get("topiclabel")
	tid := c.Input().Get("tid")

	beego.Trace("Topic Id:" + tid)

	_, fh, err := c.GetFile("attachment")
	if err != nil {
		beego.Error(err)
	}

	var attachment string
	if fh != nil {
		attachment = fh.Filename
		beego.Trace(attachment)
		err = c.SaveToFile("attachment", path.Join("attachment", attachment))
		if err != nil {
			beego.Error(err)
		}
	}

	if len(tid) == 0 {
		err = models.AddTopic(title, content, label, topicCategory, attachment)
	} else {
		err = models.ModifyTopic(tid, title, content, label, topicCategory, attachment)
	}

	if err != nil {
		beego.Error(err)
	}

	c.Redirect("/topic", 302)
}

func (c *TopicController) Add() {
	c.TplName = "topic_add.html"
	c.Data["IsLogined"] = checkAccount(c.Ctx)
}

func (c *TopicController) View() {
	c.TplName = "topic_view.html"
	c.Data["IsLogined"] = checkAccount(c.Ctx)

	reqUrl := c.Ctx.Request.RequestURI
	i := strings.LastIndex(reqUrl, "/")
	tid := reqUrl[i+1:]

	/*	params := c.Ctx.Input.Params()
		tid := params["0"]*/
	topic, err := models.GetOneTopic(tid)

	if err != nil {
		beego.Error(err)
		c.Redirect("/topic", 302)
		return
	}
	c.Data["Topic"] = topic
	c.Data["Labels"] = strings.Split(topic.Labels, " ")

	c.Data["Tid"] = tid

	c.Data["Replies"], err = models.GetAllReplies(tid)

	if err != nil {
		beego.Error(err)
		return
	}
}

func (c *TopicController) Modify() {
	beego.Trace("Modify")
	c.TplName = "topic_modify.html"
	tid := c.Input().Get("tid")
	topic, err := models.GetOneTopic(tid)
	if err != nil {
		beego.Error(err)
		c.Redirect("/topic", 302)
	}

	c.Data["Topic"] = topic
	c.Data["Tid"] = tid
}

func (c *TopicController) Delete() {
	beego.Trace("Delete")
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}

	params := c.Ctx.Input.Params()
	tid := params["0"]
	beego.Trace(tid)
	err := models.DeleteOneTopic(tid)
	if err != nil {
		beego.Error(err)

	}

	c.Redirect("/topic", 302)
}
