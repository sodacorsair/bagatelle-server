package models

import (
	"bagatelle-server/utils"
	"fmt"
	"log"
	"time"
)

type User struct {
	Id         int       `xorm:"int(11) autoincr pk"	json:"userid"`
	Username   string    `xorm:"varchar(50) notnull unique"	json:"username"`
	Password   string    `xorm:"varchar(255) notnull"	json:"password"`
	Permission int8      `xorm:"tinyint(4) default(1)"	json:"permission"`
	CreatedAt  time.Time `xorm:"created default(null)"`
	UpdatedAt  time.Time `xorm:"updated default(null)"`
}

func InsertUser(user *User) error {
	var err error
	if DB != nil {
		_, err = DB.Insert(user)
	} else {
		log.Println("insert user error: ")
		utils.ResponseError(err)
	}
	return err
}

func FindUser(user *User) (bool, error) {
	var err error
	var isExist bool
	if DB != nil {
		isExist, err = DB.Get(user)
		return isExist, err
	}
	return false, fmt.Errorf("no database existed")
}

func GetUserRows() int {
	return GetRows("user")
}