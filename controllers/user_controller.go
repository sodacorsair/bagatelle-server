package controllers

import (
	"bagatelle-server/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Login() {

}

func (c *UserController) Register() {
	var user models.User
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &user); err == nil {
		log.Printf("user: %v", user)
		user.Permission = 1
		models.InsertUser(&user)
	}

	models.FindUser(&user)
	log.Printf("user found: %v", user)
	var res map[string]interface{}
	res = map[string]interface{}{"code": 200, "userid": user.Id, "username": user.Username, "permission": user.Permission}
	log.Printf("json return %s", res)
	//jsonBody, err := json.Marshal(res)
	//if err != nil {
	//	utils.ResponseError(err)
	//}
	//log.Printf("user json: %s", jsonBody)
	c.Data["json"] = res
	c.ServeJSON()
}
