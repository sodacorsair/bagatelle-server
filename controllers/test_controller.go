package controllers

import (
	"github.com/astaxie/beego"
)

var (
	Articles map[string]*Article
)

type Article struct {
	ArticleId   string
	Title       string
	PublishTime string
	Content     string
}

func init() {
	Articles = make(map[string]*Article)
	Articles["1"] = &Article{"1", "Hello", "2016-11-08 12:00:00", "This is my first blog!"}
	Articles["2"] = &Article{"2", "Another test", "2016-11-08 13:00:00", "This is my second blog!"}
}

type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {
	ob := GetAll()
	c.Data["json"] = ob
	c.ServeJSON()
}

func GetAll() []Article {
	a := make([]Article, 0, len(Articles))
	for _, v := range Articles {
		a = append(a, *v)
	}
	return a
}
