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
	beego.Router("/articles/all", &controllers.ArticleController{}, "get:ArticlesRetrieve")
	beego.Router("/register", &controllers.UserController{}, "post:Register")
	beego.Router("/login", &controllers.UserController{}, "post:Login")
	beego.Router("/article/post", &controllers.ArticleController{}, "post:ArticlePost")
	beego.Router("/article/update", &controllers.ArticleController{}, "post:ArticleUpdate")
	beego.Router("/article/get/?:id", &controllers.ArticleController{}, "get:ArticleRetrieve")
	beego.Router("/tags/all", &controllers.TagController{}, "get:TagsRetrieve")
	beego.Router("/tag/get", &controllers.TagController{}, "get:ArticlesRetrieveByTag")
	beego.Router("/categories/all", &controllers.CategoryController{}, "get:CatesRetrieve")
	beego.Router("/category/get", &controllers.CategoryController{}, "get:ArticlesRetrieveByCate")
	beego.Router("/articles/manage", &controllers.ArticleController{}, "get:ArticlesManage")
}
