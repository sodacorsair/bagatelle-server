package controllers

import (
	"bagatelle-server/models"
	"bagatelle-server/utils"
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Register() {
	var user models.User
	var res map[string]interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &user); err == nil {
		log.Printf("user: %v", user)
		if isExist, _ := models.FindUser(&models.User{Username: user.Username}); isExist {
			res = map[string]interface{}{"code": 400, "message": "用户名已被注册!"}
		} else {
			//salt1 := "11"
			//salt2 := "22"
			//h := md5.New()
			//io.WriteString(h, user.Password)
            //pwdm5 := fmt.Sprintf("%x", h.Sum(nil))
			//io.WriteString(h, salt1)
			//io.WriteString(h, user.Username)
			//io.WriteString(h, salt2)
			//io.WriteString(h, pwdm5)

            //user.Password = fmt.Sprintf("%x", h.Sum(nil))

			user.Password = utils.CryptPwd(user.Username, user.Password)

			user.Permission = 1

			models.InsertUser(&user)

			isExist, _ = models.FindUser(&user)
			log.Println("user existed: ", isExist)
			log.Printf("user found: %v", user)
			log.Printf("password: %s", user.Password)
			res = map[string]interface{}{"code": 200, "userid": user.Id, "username": user.Username, "permission": user.Permission}
		}
	} else {
		res = map[string]interface{}{"code": 400, "message": "网络错误"}
	}

	log.Printf("json return %s", res)
	c.Data["json"] = res
	c.ServeJSON()
}

func (c *UserController) Login() {
    var user models.User
    var res map[string]interface{}

    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &user); err != nil {
    	utils.ResponseError(err)
    	res = map[string]interface{}{"code": 400, "message": "网络错误"}
	} else if isExist, _ := models.FindUser(&models.User{Username: user.Username}); !isExist {
        res = map[string]interface{}{"code": 400, "message": "用户名不存在"}
	} else if isExist, _ := models.FindUser(&models.User{Username: user.Username, Password: utils.CryptPwd(user.Username, user.Password)}); !isExist {
		res = map[string]interface{}{"code": 400, "message": "密码不正确"}
	} else {
		models.FindUser(&user)
		res = map[string]interface{}{"code": 200, "userid": user.Id, "username": user.Username, "permission": user.Permission}
	}

	c.Data["json"] = res
	c.ServeJSON()
}
