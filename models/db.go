package models

import (
	"bagatelle-server/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
)

type User struct {
	Id         int       `xorm:"int(11) notnull autoincr"`
	Username   string    `xorm:"varchar(50) notnull unique"`
	Password   string    `xorm:"varchar(255) notnull"`
	Permission int8      `xorm:"tinyint(4) default(0)"`
	CreatedAt  time.Time `xorm:"created default(null)"`
	UpdatedAt  time.Time `xorm:"updated default(null)"`
}

type Article struct {
	Id        int       `xorm:"int(11) notnull autoincr"`
	Title     string    `xorm:"notnull"`
	Content   string    `xorm:"text"`
	CreatedAt time.Time `xorm:"created default(null)"`
	UpdatedAt time.Time `xorm:"updated default(null)"`
}

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
	db *xorm.Engine
)

func InitDB() {
	db, err := xorm.NewEngine("mysql", "root:225500@tcp(127.0.0.1:3306)/bagatelle?charset=utf8")
	utils.ResponseError(err)

	//createTables(db, User{}, Article{}, Comment{}, Category{}, Tag{})
	if isExist, _ := db.IsTableExist(&User{}); !isExist {
		err := db.CreateTables(&User{})
		utils.ResponseError(err)
	}
	if isExist, _ := db.IsTableExist(&Article{}); !isExist {
		err := db.CreateTables(&Article{})
		utils.ResponseError(err)
	}
	if isExist, _ := db.IsTableExist(&Comment{}); !isExist {
		err := db.CreateTables(&Comment{})
		utils.ResponseError(err)
	}
	if isExist, _ := db.IsTableExist(&Category{}); !isExist {
		err := db.CreateTables(&Category{})
		utils.ResponseError(err)
	}
	if isExist, _ := db.IsTableExist(&Tag{}); !isExist {
		err := db.CreateTables(&Tag{})
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
