package models

import (
	"bagatelle-server/utils"
	"errors"
	"strconv"
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

func DeleteCategories(article Article) {
	if DB != nil {
		var cates []Category
		FindCategories(&cates, "article_id=" + strconv.Itoa(article.Id))
		for _, cate := range cates {
			DB.Id(cate.Id).Delete(cate)
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

func UpdateCategories(article Article, updateCates []string) {
	if DB != nil {
		var cates []Category
		FindCategories(&cates, "article_id=" + strconv.Itoa(article.Id))
		for i := 0; i < len(cates); i++ {
			DB.Id(cates[i].Id).Delete(cates[i])
		}
		for _, c := range updateCates {
			cate := Category{
				Name:      c,
				ArticleId: article.Id,
			}
			InsertCategory(&cate)
		}
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
