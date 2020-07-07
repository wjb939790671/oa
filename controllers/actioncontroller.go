package controllers

import (
	"OA/bll"
	"OA/models"
	"OA/utils"
)

type ActionController struct {
	BaseController
}

func (this *ActionController) Prepare() {
	this.model = &models.Action{}
	var list []models.Action
	this.list = &list
	this.iBll = &bll.ActionBll{}
	this.pathFolder = "action/"
<<<<<<< HEAD
	this.unique = "Url"
=======
	this.admin = true

}

func (this *ActionController) QueryAll() {
	where := make(map[string]interface{})
	if !this.admin {
		where["Delflage"] = false
	}
	where["IsMenu"] = true
	this.iBll.QueryAll(this.model, this.list, where, []string{"Id", "ActionName"}, "Id", "ActionName")
	this.message.Code = utils.OK
	this.message.Text = utils.GetCodeText(this.message.Code)
	this.message.Data = this.list
	this.Data["json"] = this.message
	this.ServeJSON()
>>>>>>> 8a1f30b2a36b34695cf2098c09d2fbb2e1fcb52f
}
