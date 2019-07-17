package routers

import (
	"bagatelle-server/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//ns := beego.NewNamespace("/v1",
	//	beego.NSNamespace("/article",
	//		beego.NSInclude(
	//			&controllers.TestController{},
	//		),
	//	),
	//)
	//beego.AddNamespace(ns)
	beego.Router("awsl", &controllers.TestController{})
	beego.Router("/register", &controllers.UserController{}, "post:Register")
	beego.Router("/login", &controllers.UserController{}, "post:Login")
	beego.Router("/article/post", &controllers.ArticleController{}, "post:ArticlePost")
	beego.Router("/article/get/?:id", &controllers.ArticleController{}, "get:ArticleRetrieve")
}
