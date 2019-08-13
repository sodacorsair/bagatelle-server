package controllers

import (
	"bagatelle-server/models"
	"bagatelle-server/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
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

func (c *ArticleController) ArticleRetrieve() {
	idStr := c.Ctx.Input.Param(":id")
	articleId, _ := strconv.Atoi(idStr)
	article := models.Article{Id: articleId}
	models.FindArticle(&article)
	tags := make([]models.Tag, 0)
	models.FindTags(&tags, "article_id="+idStr)
	for _, t := range tags {
		fmt.Printf("%s\n", t.Name)
	}
	cates := make([]models.Category, 0)
	models.FindCategories(&cates, "article_id="+idStr)
	for _, c := range cates {
		fmt.Printf("category: %s\n", c.Name)
	}


    postTime := utils.TimeFormat(article.CreatedAt)
    updateTime := utils.TimeFormat(article.UpdatedAt)

	res := map[string]interface{}{
		"code":       200,
		"title":      article.Title,
		"content":    article.Content,
		"isPrivate":  article.Private,
		"tags":       tags,
		"cates":      cates,
		"postTime":   postTime,
		"updateTime": updateTime,
		"reads":      article.Reads,
	}
	c.Data["json"] = res
	c.ServeJSON()
}

func (c *ArticleController) ArticlesRetrieve() {
	type Item struct {
		Id int `json:"id"`
		Name string `json:"name"`
		CreateAt string `json:"createdAt"`
	}

	articles := make([]models.Article, 0)
	models.FindAllArticles(&articles)
	articleList := make([]Item, len(articles))

	for i := 0; i < len(articles); i++ {
		articleList[i].Name = articles[i].Title
		articleList[i].Id = articles[i].Id
		articleList[i].CreateAt = utils.ShortTimeFormat(articles[i].CreatedAt)
	}

	res := map[string]interface{}{
		"code": 200,
		"articles": articleList,
	}

	c.Data["json"] = res
	c.ServeJSON()
}