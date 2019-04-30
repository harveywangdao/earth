package routers

import (
	"github.com/astaxie/beego"
	"os"
	"zeus/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{})
	beego.Router("/reply", &controllers.ReplyController{})
	beego.Router("/reply/add", &controllers.ReplyController{}, "post:Add")
	beego.Router("/reply/delete", &controllers.ReplyController{}, "get:Delete")

	os.Mkdir("attachment", os.ModePerm)
	//beego.SetStaticPath("/attachment", "attachment")
	beego.Router("/attachment/:all", &controllers.AttachController{})
}
