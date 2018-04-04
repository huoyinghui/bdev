package controllers

import (
	"bjwt/models"
	"bjwt/utils"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about Token
type TokenController struct {
	beego.Controller
}

// @Title Create
// @Description create token
// @Param	body		body 	models.UserAuth	true		"The UserAuth content"
// @Success 200 {string} models.UserAuth.Uid
// @Success 200 {string} models.UserAuth.Secret
// @Failure 403 body is empty
// @router / [post]
func (o *TokenController) Post() {
	var ob models.UserAuth
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	secret := utils.SecSecret(ob.Uid, "123")
	token, err := utils.CreateToken(ob.Uid, secret, 0)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = map[string]string{
			"token": token,
			//"secret": secret,
		}
	}
	o.ServeJSON()
}

func (o *TokenController) Option() {
	o.Data["json"] = ""
	o.ServeJSON()
}
