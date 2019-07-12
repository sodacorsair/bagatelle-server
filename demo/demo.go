package main

import (
	"encoding/json"
	"log"
)

func main() {
	//db, _ := xorm.NewEngine("mysql", "root:225500@tcp(127.0.0.1:3306)/demo?charset=utf8")
	//db.CreateTables(&models.User{})
	//user := models.User{Username: "a", Password: "a"}
	////err := models.InsertUser("a", "a")
	//models.DB, _ = xorm.NewEngine("mysql", "root:225500@tcp(127.0.0.1:3306)/bagatelle?charset=utf8")
	//var err error
	//if models.DB != nil {
	//	_, err = models.DB.Insert(&user)
	//}
	//utils.ResponseError(err)

	type Person struct {
		Name string
		Age  int
	}
	person := Person{"shabi", 22}
	jsonBytes, _ := json.Marshal(person)
	log.Printf("%s", jsonBytes)
}
