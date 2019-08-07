package models

import (
	"bagatelle-server/utils"
	"errors"
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

func FindAllTags(tags *[]Tag) {
	if DB != nil {
		err := DB.Table("tag").Find(tags)
		if err != nil {}
		utils.ResponseError(err)
	} else {
		utils.ResponseError(errors.New("DB not existed"))
	}
}
