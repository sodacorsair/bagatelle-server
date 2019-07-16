package models

import (
	"bagatelle-server/utils"
	"errors"
)

type Tag struct {
	Id        int    `xorm:"int(11) autoincr pk"`
	Name      string `xorm:"varchar(100) notnull"`
	ArticleId int    `xorm:"int(11) default(null)"`
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
