package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/beego/i18n"
	"zeus/models"
	_ "zeus/routers"
)

func init() {
	models.RegisterDB()
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	i18n.SetMessage("zh-CN", "conf/"+"locale_"+"zh-CN"+".ini")
	i18n.SetMessage("en-US", "conf/"+"locale_"+"en-US"+".ini")
	beego.AddFuncMap("i18n", i18n.Tr)

	beego.Run()
}
