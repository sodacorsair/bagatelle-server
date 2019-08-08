package controllers

import (
	"bagatelle-server/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

type TagController struct {
	beego.Controller
}

func (c *TagController) TagsRetrieve() {
	tags := make([]models.Tag, 0)
	models.FindAllTags(&tags)

	res := map[string]interface{}{
		"code": 200,
		"taglist": tags,
	}
	c.Data["json"] = res
	c.ServeJSON()
}

func (c *TagController) ArticlesRetrieveByTag() {
	type Params struct {
		name string
		page int
		pageSize int
	}
	var params Params

	json.Unmarshal(c.Ctx.Input.RequestBody, &params)

	name := params.name

	var tags *[]models.Tag
	models.FindTags(tags, "name=" + name)

	articles := make([]models.Article, len(*tags))
	for i := 0; i < len(*tags); i++ {
		articles[i].Id = (*tags)[i].ArticleId
		models.FindArticle(&articles[i])
	}

	res := map[string]interface{}{
		"code": 200,
		"articlelist": articles,
	}

	c.Data["json"] = res
	c.ServeJSON()

}
