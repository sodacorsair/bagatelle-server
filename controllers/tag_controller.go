package controllers

import (
	"bagatelle-server/models"
	"bagatelle-server/utils"
	"github.com/astaxie/beego"
)

type TagController struct {
	beego.Controller
}

type Item struct {
	Name string	`json:"name"`
	Id int	`json:"id"`
	CreatedAt string `json:"createdAt"`
}

type ItemArr []Item

func (arr ItemArr) Len() int {
	return len(arr)
}

func (arr ItemArr) Swap(i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func (c *TagController) TagsRetrieve() {
	tags := make([]models.Tag, 0)
	models.FindAllTags(&tags)

	res := map[string]interface{}{
		"code": 200,
		"list": tags,
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

	var items ItemArr
	items = make([]Item, len(articles))
	for i := 0; i < len(articles); i++ {
		items[i].Name = articles[i].Title
		items[i].Id = articles[i].Id
		items[i].CreatedAt = utils.ShortTimeFormat(articles[i].CreatedAt)
	}

	utils.ReverseArray(items)

	res := map[string]interface{}{
		"code": 200,
		"articleItems": items,
	}

	c.Data["json"] = res
	c.ServeJSON()

}
