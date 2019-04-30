package controllers

import (
	beego "github.com/astaxie/beego"
	"github.com/beego/i18n"
	"zeus/models"
)

type baseController struct {
	beego.Controller
	i18n.Locale
}

func (c *baseController) Prepare() {
	lang := c.GetString("lang")
	if lang == "zh-CN" {
		c.Lang = lang
	} else {
		c.Lang = "en-US"
	}

	beego.Trace("Prepare")

	c.Data["Lang"] = c.Lang
}

type MainController struct {
	baseController
}

func (c *MainController) Get() {
	c.Data["IsHome"] = true
	c.TplName = "home.html"
	c.Data["IsLogined"] = checkAccount(c.Ctx)

	c.Data["TopicCategory"] = "topic_category"

	cate := c.Input().Get("cate")
	label := c.Input().Get("label")

	var err error
	c.Data["Topics"], err = models.GetAllTopics(cate, label, true)
	if err != nil {
		beego.Error(err)
	}

	c.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
}
