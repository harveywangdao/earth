package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	isExit := c.Input().Get("exitLogin") == "true"
	if isExit {
		c.Ctx.SetCookie("account", "", -1, "/")
		c.Ctx.SetCookie("password", "", -1, "/")
		c.Redirect("/", 301)
		return
	}

	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	act := c.Input().Get("account")
	pswd := c.Input().Get("password")
	aulogin := c.Input().Get("autologin") == "on"

	if beego.AppConfig.String("account") == act &&
		beego.AppConfig.String("password") == pswd {
		maxAccount := 0
		if aulogin {
			maxAccount = 1<<32 - 1
		}
		c.Ctx.SetCookie("account", act, maxAccount, "/")
		c.Ctx.SetCookie("password", pswd, maxAccount, "/")
	}

	c.Redirect("/", 301)
	return
}

func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("account")
	if err != nil {
		return false
	}

	act := ck.Value

	ck, err = ctx.Request.Cookie("password")
	if err != nil {
		return false
	}

	pswd := ck.Value
	beego.Trace(act, pswd)
	return beego.AppConfig.String("account") == act &&
		beego.AppConfig.String("password") == pswd
}
