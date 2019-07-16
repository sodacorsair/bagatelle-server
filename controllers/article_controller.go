package controllers

import (
	"bagatelle-server/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) ArticlePost() {
	type PostReceived struct {
		Tags    []string `json:"submitTags"`
		Cates   []string `json:"submitCates"`
		Title   string   `json:"submitTitle"`
		Content string   `json:"submitContent"`
		Private bool     `json:"submitPrivate"`
		Top     bool     `json:"submitTop"`
	}
	var postReceived PostReceived
	var res map[string]interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &postReceived); err == nil {
		article := models.Article{
			Title:   postReceived.Title,
			Content: postReceived.Content,
			Private: postReceived.Private,
			Top:     postReceived.Top,
		}
		models.InsertArticle(&article)
		models.FindArticle(&article)
		for _, t := range postReceived.Tags {
			tag := models.Tag{
				Name:      t,
				ArticleId: article.Id,
			}
			models.InsertTag(&tag)
		}
		for _, c := range postReceived.Cates {
			cate := models.Category{
				Name:      c,
				ArticleId: article.Id,
			}
			models.InsertCategory(&cate)
		}
		res = map[string]interface{}{"code": 200, "articleid": article.Id}
	} else {
		res = map[string]interface{}{"code": 400, "message": "网络错误"}
	}

	c.Data["json"] = res
	c.ServeJSON()
}
