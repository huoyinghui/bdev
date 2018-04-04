package controllers

import (
	"bdev/config"
	"bdev/models"
	"bdev/utils"
	"encoding/json"

	"github.com/astaxie/beego"
	"bdev/logger"
)

// Operations about Token
type ValidController struct {
	beego.Controller
}

// @Title Valid
// @Description valid token
// @Param	body		body 	models.UserAuth	true		"The UserAuth content"
// @Success 200 {string} models.UserAuth.Token
// @Success 200 {string} models.UserAuth.Secret
// @Failure 403 body is empty
// @router / [post]
func (o *ValidController) Post() {
	var ob models.UserAuth
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	uid, err := utils.GetUid(ob.Token)
	if err != nil {
		logger.Errorf("GetUid: %v", err)
		o.Data["json"] = err.Error()
		o.ServeJSON()
		return
	}
	secret := utils.SecSecret(uid, config.AppConf.JwtSalt)
	logger.Debugf("sec:%v", secret)
	newUid, err := utils.AuthToken(ob.Token, secret)
	if err != nil {
		logger.Errorf("AuthToken: %v", err)
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = map[string]string{"uid": uid, "new_uid": newUid}
	}
	o.ServeJSON()
}

func (o *ValidController) Option() {
	o.Data["json"] = ""
	o.ServeJSON()
}
