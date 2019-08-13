package models

import (
	"bagatelle-server/utils"
	"errors"
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

func FindAllArticles(articles *[]Article) {
	if DB != nil {
		err := DB.Table("article").Find(articles)
		if err != nil {}
		utils.ResponseError(err)
	} else {
		utils.ResponseError(errors.New("DB not existed"))
	}
}