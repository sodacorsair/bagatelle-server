package controllers

import (
	"bagatelle-server/models"
	"fmt"
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
	}
	items := make([]Item, len(articles))
	for i := 0; i < len(articles); i++ {
		items[i].Name = articles[i].Title
		items[i].Id = articles[i].Id
	}

	for _, item := range items {
		fmt.Printf("%v\n", item)
	}

	res := map[string]interface{}{
		"code": 200,
		"articlelist": items,
	}

	c.Data["json"] = res
	c.ServeJSON()

}
