package models

import (
	"bagatelle-server/utils"
	"errors"
	"strconv"
	"time"
)

type Article struct {
	Id        int       `xorm:"int(11) autoincr pk"`
	Title     string    `xorm:"notnull" json:"submitTitle"`
	Content   string    `xorm:"text" json:"submitContent"`
	Top       bool      `json:"submitTop"`
	Private   bool      `json:"submitPrivate"`
	Reads     int       `xorm:"default(0)"`
	CreatedAt time.Time `xorm:"created default(null)"`
	UpdatedAt time.Time `xorm:"updated default(null)"`
}

func GetArticleRows() int {
	return GetRows("article")
}

func InsertArticle(article *Article) {
	if DB != nil {
		_, err := DB.Insert(article)
		if err != nil {
			utils.ResponseError(err)
		}
	} else {
		utils.ResponseError(errors.New("DB not existed"))
	}
}

func DeleteArticle(article *Article) {
	if DB != nil {
		DB.Id(article.Id).Delete(article)
	} else {
		utils.ResponseError(errors.New("DB really not existed!"))
	}
}

func UpdateArticle(article *Article) {
	if DB != nil {
		_, err := DB.Id(article.Id).Update(article)
		if err != nil {
			utils.ResponseError(err)
		}
	} else {
		utils.ResponseError(errors.New("DB not existed"))
	}
}

func FindArticle(article *Article) bool {
	if DB != nil {
		isExist, err := DB.Get(article)
		if err == nil {
			return isExist
		} else {
			utils.ResponseError(err)
		}
	} else {
		utils.ResponseError(errors.New("DB not existed"))
	}
	return false
}

func FindArticles(articles *[]Article, b, a int) {
	if DB != nil {
		result, _ := DB.Query("select count(*) from article")
		newStr := string(result[0]["count(*)"])
		rows, _ := strconv.Atoi(newStr)
		err := DB.Table("article").Limit(a - b + 1, rows - a).Desc("id").Find(articles)
		if err != nil {
			utils.ResponseError(err)
		}
	} else {
		utils.ResponseError(errors.New("DB not existed"))
	}
}