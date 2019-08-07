package controllers

import (
	"bagatelle-server/models"
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
