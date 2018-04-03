package controllers

import (
	"bdev/models"
	"encoding/json"

	"bdev/logger"
	"fmt"
	"github.com/astaxie/beego"
)

// Operations about object
type MessageController struct {
	beego.Controller
}

// @Title Create
// @Description create object
// @Param	body		body 	models.Message	true		"The Message content"
// @Success 200 {string} models.Message.Id
// @Failure 403 body is empty
// @router / [post]
func (o *MessageController) CreateMessage() {
	var ob models.MessageJson
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	if err != nil {
		logger.Errorf("err:%v", err)
		o.Abort("500")
	}
	logger.Debugf("argJson:%v", ob)
	objectid, err := models.AddMessage(&models.Message{
		Id:         ob.Id,
		FromUserId: ob.FromUserId,
		ToUserId:   ob.ToUserId,
		Title:      ob.Title,
		Message:    ob.Message,
	})
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = map[string]interface{}{"ObjectId": objectid}
	}
	o.ServeJSON()
}

func (o *MessageController) ReadMessage() {
	var ob models.MessageJson
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	if err != nil {
		logger.Errorf("err:%v", err)
		o.Abort("500")
	}
	logger.Debugf("argJson:%v", ob)

	nums := int64(0)
	msgList := []*models.Message{}
	// todo
	//models.ReadMessage()
	if ob.FromUserId > 0 {
		nums, msgList, err = models.ReadMessageWithFromUserId(ob.MsgType, ob.Id, ob.ToUserId, ob.FromUserId, ob.PageNumber, ob.PageSize)
	} else {
		nums, msgList, err = models.ReadMessage(ob.MsgType, ob.Id, ob.ToUserId, ob.PageNumber, ob.PageSize)
	}
	if err != nil {
		logger.Errorf("ShowToUserMessages err:%v argJson:%v", err, ob)
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = map[string]interface{}{
			"err":  "",
			"nums": nums,
			"data": msgList,
			"pageNumber": ob.PageNumber,
			"pageSize": ob.PageSize,
		}
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the msg
// @Param	Id		path 	string	true		"The message Id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (o *MessageController) DeleteMessage() {
	var ob models.MessageJson
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	if err != nil {
		logger.Errorf("err:%v", err)
		o.Abort("500")
	}
	logger.Debugf("argJson:%v", ob)
	if err := models.DeleteMessage(ob.Id); err != nil {
		logger.Errorf("err:%v", err)
		o.Data["json"] = fmt.Errorf("DeleteMessage src err:%v ob.id:%v", err, ob.Id)
	} else {
		o.Data["json"] = "delete success!"
	}
	o.ServeJSON()
}

func (o *MessageController) UpdateMessage() {
	var ob models.MessageJson
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	if err != nil {
		logger.Errorf("err:%v", err)
		o.Abort("500")
	}
	logger.Debugf("argJson:%v", ob)
	msg, err := models.UpdateMessage(ob.Id, &models.Message{
		Id:     ob.Id,
		Status: ob.Status,
	})
	if err != nil {
		logger.Errorf("UpdateMessage err:%v argJson:%v", err, ob)
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = msg
	}
	o.ServeJSON()
}
