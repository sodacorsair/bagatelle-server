package models

import (
	"bagatelle-server/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
)

type Category struct {
	Id        int    `xorm:"int(11) notnull autoincr"`
	Name      string `xorm:"varchar(100) notnull"`
	ArticleId int    `xorm:"int(11) default(null)"`
}

type Tag struct {
	Id        int    `xorm:"int(11) notnull autoincr"`
	Name      string `xorm:"varchar(100) notnull"`
	ArticleId int    `xorm:"int(11) default(null)"`
}

type Comment struct {
	Id        int       `xorm:"int(11) notnull autoincr"`
	UserId    int       `xorm:"int(11) default(null)"`
	ArticleId int       `xorm:"int(11) default(null)"`
	Content   string    `xorm:"text"`
	CreatedAt time.Time `xorm:"created default(null)"`
	UpdatedAt time.Time `xorm:"updated default(null)"`
}

var (
	DB *xorm.Engine
)

func InitDB() {
	var err error
	DB, err = xorm.NewEngine("mysql", "root:225500@tcp(127.0.0.1:3306)/bagatelle?charset=utf8")
	utils.ResponseError(err)

	//createTables(DB, User{}, Article{}, Comment{}, Category{}, Tag{})
	if isExist, _ := DB.IsTableExist(&User{}); !isExist {
		err := DB.CreateTables(&User{})
		utils.ResponseError(err)
	}
	if isExist, _ := DB.IsTableExist(&Article{}); !isExist {
		err := DB.CreateTables(&Article{})
		utils.ResponseError(err)
	}
	if isExist, _ := DB.IsTableExist(&Comment{}); !isExist {
		err := DB.CreateTables(&Comment{})
		utils.ResponseError(err)
	}
	if isExist, _ := DB.IsTableExist(&Category{}); !isExist {
		err := DB.CreateTables(&Category{})
		utils.ResponseError(err)
	}
	if isExist, _ := DB.IsTableExist(&Tag{}); !isExist {
		err := DB.CreateTables(&Tag{})
		utils.ResponseError(err)
	}
}

func createTable(engine *xorm.Engine, tableName interface{}) {
	if isExist, _ := engine.IsTableExist(&tableName); !isExist {
		tableName = tableName.(Article)
		err := engine.CreateTables(&tableName)
		utils.ResponseError(err)
	}
}

func createTables(engine *xorm.Engine, tableNames ...interface{}) {
	for _, tableName := range tableNames {
		createTable(engine, tableName)
	}
}
