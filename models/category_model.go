package models

import (
	"bagatelle-server/utils"
	"errors"
)

type Category struct {
	Id        int    `xorm:"int(11) autoincr pk"	json:"id"`
	Name      string `xorm:"varchar(100) notnull"	json:"name"`
	ArticleId int    `xorm:"int(11) default(null)"	json:"article_id"`
}

func InsertCategory(category *Category) {
	if DB != nil {
		_, err := DB.Insert(category)
		if err != nil {
			utils.ResponseError(err)
		}
	} else {
		utils.ResponseError(errors.New("DB not existed"))
	}
}

func FindCategories(cates *[]Category, sql string) {
	if DB != nil {
		err := DB.Table("category").Where(sql).Find(cates)
		if err != nil {
			utils.ResponseError(err)
		}
	} else {
		utils.ResponseError(errors.New("DB not existed"))
	}
}

func FindAllCategories(cates *[]Category) {
	if DB != nil {
		err := DB.Table("category").Find(cates)
		if err != nil {}
		utils.ResponseError(err)
	} else {
		utils.ResponseError(errors.New("DB not existed"))
	}
}
