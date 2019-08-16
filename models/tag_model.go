package models

import (
	"bagatelle-server/utils"
	"errors"
	"strconv"
)

type Tag struct {
	Id        int    `xorm:"int(11) autoincr pk"	json:"id"`
	Name      string `xorm:"varchar(100) notnull"	json:"name"`
	ArticleId int    `xorm:"int(11) default(null)"	json:"article_id"`
}

func InsertTag(tag *Tag) {
	if DB != nil {
		_, err := DB.Insert(tag)
		if err != nil {
			utils.ResponseError(err)
		}
	} else {
		utils.ResponseError(errors.New("DB not existed"))
	}
}

func FindTags(tags *[]Tag, sql string) {
	if DB != nil {
		err := DB.Table("tag").Where(sql).Find(tags)
		if err != nil {
			utils.ResponseError(err)
		}
	} else {
		utils.ResponseError(errors.New("DB not existed"))
	}
}

func UpdateTags(article Article, updatedTags []string) {
	if DB != nil {
		var tags []Tag
		FindTags(&tags, "article_id=" + strconv.Itoa(article.Id))
		for i := 0; i < len(tags); i++ {
			DB.Id(tags[i].Id).Delete(tags[i])
		}
		for _, t := range updatedTags {
			tag := Tag{
				Name:      t,
				ArticleId: article.Id,
			}
			InsertTag(&tag)
		}
	}
}

func FindAllTags(tags *[]Tag) {
	if DB != nil {
		err := DB.Table("tag").Find(tags)
		if err != nil {}
		utils.ResponseError(err)
	} else {
		utils.ResponseError(errors.New("DB not existed"))
	}
}
