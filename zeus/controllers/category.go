package controllers

import (
	"github.com/astaxie/beego"
	"zeus/models"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	op := c.Input().Get("op")
	beego.Trace("option:" + op)
	switch op {
	case "add":
		name := c.Input().Get("categoryname")
		if len(name) == 0 {
			break
		}
		beego.Trace("Category Title:" + name)
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/category", 301)
		return
	case "delete":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		beego.Trace("Category Id:", id)
		err := models.DeleteCategory(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/category", 301)
		return
	}

	c.Data["IsCategory"] = true
	c.TplName = "category.html"
	c.Data["IsLogined"] = checkAccount(c.Ctx)

	var err error
	c.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
}
