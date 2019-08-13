package controllers

import (
	"bagatelle-server/models"
	"bagatelle-server/utils"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) CatesRetrieve() {
	cates := make([]models.Category, 0)
	models.FindAllCategories(&cates)

	res := map[string]interface{}{
		"code": 200,
		"catelist": cates,
	}
	c.Data["json"] = res
	c.ServeJSON()
}

func (c *CategoryController) ArticlesRetrieveByCate() {
	name := c.GetString("name")

	cates := make([]models.Category, 0)
	models.FindCategories(&cates, "name='" + name + "'")

	articles := make([]models.Article, len(cates))
	for i := 0; i < len(cates); i++ {
		articles[i].Id = cates[i].ArticleId
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
