package controllers

import (
	"bagatelle-server/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) ArticlePost() {
	var article models.Article
	var res map[string]interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &article); err == nil {
		log.Printf("article: %v", article)
		models.InsertArticle(&article)
		models.FindArticle(&article)
		res = map[string]interface{}{"code": 200, "articleid": article.Id}
	} else {
		res = map[string]interface{}{"code": 400, "message": "网络错误"}
	}

	c.Data["json"] = res
	c.ServeJSON()
}
