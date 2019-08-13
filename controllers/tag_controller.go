package controllers

import (
	"bagatelle-server/models"
	"bagatelle-server/utils"
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
	name := c.GetString("name")

	tags := make([]models.Tag, 0)
	models.FindTags(&tags, "name='" + name + "'")

	articles := make([]models.Article, len(tags))
	for i := 0; i < len(tags); i++ {
		articles[i].Id = tags[i].ArticleId
		models.FindArticle(&(articles[i]))
	}

	type Item struct {
		Name string	`json:"name"`
		Id int	`json:"id"`
		CreatedAt string `json:"createdAt"`
	}
	items := make([]Item, len(articles))
	for i := 0; i < len(articles); i++ {
		items[i].Name = articles[i].Title
		items[i].Id = articles[i].Id
		items[i].CreatedAt = utils.ShortTimeFormat(articles[i].CreatedAt)
	}

	res := map[string]interface{}{
		"code": 200,
		"articleItems": items,
	}

	c.Data["json"] = res
	c.ServeJSON()

}
