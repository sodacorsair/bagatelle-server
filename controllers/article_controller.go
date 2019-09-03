package controllers

import (
	"bagatelle-server/models"
	"bagatelle-server/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"log"
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

func (c *ArticleController) ArticleUpdate() {
	type PostReceived struct {
		Tags    []string `json:"submitTags"`
		Cates   []string `json:"submitCates"`
		Title   string   `json:"submitTitle"`
		Content string   `json:"submitContent"`
		Private bool     `json:"submitPrivate"`
		Top     bool     `json:"submitTop"`
		Id	string	`json:"submitId"`
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
		article.Id, _ = strconv.Atoi(postReceived.Id)
		models.UpdateArticle(&article)
		models.FindArticle(&article)
		models.UpdateCategories(article, postReceived.Cates)
		models.UpdateTags(article, postReceived.Tags)

		res = map[string]interface{}{"code": 200, "articleid": article.Id}
	} else {
		res = map[string]interface{}{"code": 400, "message": "网络错误"}
	}

	c.Data["json"] = res
	c.ServeJSON()
}

func (c *ArticleController) ArticleDelete() {
	idStr := c.Ctx.Input.Param(":id")
	articleId, _ := strconv.Atoi(idStr)
	article := models.Article{Id: articleId}
	models.FindArticle(&article)
	models.DeleteCategories(article)
	models.DeleteTags(article)
	models.DeleteArticle(&article)
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

	article.Reads++
	models.UpdateArticle(&article)

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

func (c *ArticleController) ArticleFirst() {
	var articles []models.Article
	models.FindArticles(&articles, 1, 1)
	article := articles[0]

	article.Reads++
	models.UpdateArticle(&article)

	res := map[string]interface{}{
		"code":       200,
		"title":      article.Title,
		"content":    article.Content,
	}
	c.Data["json"] = res
	c.ServeJSON()
}

func (c *ArticleController) ArticlesRetrieve() {
	type ArticleItem struct {
		Id int `json:"id"`
		Name string `json:"name"`
		CreatedAt string `json:"createdAt"`
	}

	page, _ := c.GetInt("page")
	pageSize, _ := c.GetInt("pageSize")

	maxId := models.GetArticleRows() - 1
	totalPages := (maxId + 1) / pageSize + 1

	var startId int
	endId := maxId - (page - 1) * pageSize + 1
	if page == totalPages {
		startId = 2
	} else {
		startId = maxId + 1 - page * pageSize + 1
	}

	log.Printf("%v\n%v\n", startId, endId)

	articles := make([]models.Article, 0)
	models.FindArticles(&articles, startId, endId)
	articleList := make([]ArticleItem, len(articles))

	for i := 0; i < len(articles); i++ {
		articleList[i].Name = articles[i].Title
		articleList[i].Id = articles[i].Id
		articleList[i].CreatedAt = utils.ShortTimeFormat(articles[i].CreatedAt)
	}

	res := map[string]interface{}{
		"code": 200,
		"articles": articleList,
		"total": maxId,
	}

	c.Data["json"] = res
	c.ServeJSON()
}

func (c *ArticleController) ArticlesManage() {
	type ArticleItem struct {
		Id int `json:"id"`
		Name string `json:"name"`
		CreatedAt string `json:"createdAt"`
		Reads int	`json:"reads"`
	}

	page, _ := c.GetInt("page")
	pageSize, _ := c.GetInt("pageSize")

	maxId := models.GetArticleRows()

	var startId int
	endId := maxId - (page - 1) * pageSize
	if page == 1 {
		startId = 1
	} else {
		startId = maxId - page * pageSize + 1
	}

	articles := make([]models.Article, 0)
	models.FindArticles(&articles, startId, endId)
	articleList := make([]ArticleItem, len(articles))

	for i := 0; i < len(articles); i++ {
		articleList[i].Name = articles[i].Title
		articleList[i].Id = articles[i].Id
		articleList[i].CreatedAt = utils.ShortTimeFormat(articles[i].CreatedAt)
		articleList[i].Reads = articles[i].Reads
	}

	res := map[string]interface{}{
		"code": 200,
		"articles": articleList,
		"total": maxId,
	}

	c.Data["json"] = res
	c.ServeJSON()
}

func (c *ArticleController) ArticlesRecent() {
	type ArticleItem struct {
		Id int `json:"id"`
		Name string `json:"name"`
		CreatedAt string `json:"createdAt"`
		Content string	`json:"content"`
	}

	maxId := models.GetArticleRows()

	pageSize := 5

	groupNum := maxId / pageSize + 1

	var startId int
	endId := maxId
	if 1 == groupNum {
		startId = 2
	} else {
		startId = maxId - pageSize + 1
	}

	articles := make([]models.Article, 0)
	models.FindArticles(&articles, startId, endId)
	articleList := make([]ArticleItem, len(articles))

	for i := 0; i < len(articles); i++ {
		articleList[i].Name = articles[i].Title
		articleList[i].Id = articles[i].Id
		articleList[i].CreatedAt = utils.ShortTimeFormat(articles[i].CreatedAt)
		articleList[i].Content = articles[i].Content
	}

	res := map[string]interface{}{
		"code": 200,
		"articles": articleList,
		"total": maxId,
	}

	c.Data["json"] = res
	c.ServeJSON()
}