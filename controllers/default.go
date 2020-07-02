package controllers

import (
	"OA/bll"
	"OA/models"
)

type MainController struct {
	BaseController
}

func (this *MainController) Prepare() {
	this.model = &models.Action{}
	var list models.Action
	this.list = &list
	this.iBll = &bll.ActionBll{}
	this.admin = true
}

func (this *MainController) Get() {
	var list []models.Action
	where := make(map[string]interface{})
	if this.admin {
		where["ParentId"] = 0
	} else {
		where["Delflage"] = false

	}
	this.iBll.QueryAll(&models.Action{}, &list, where, []string{"Id", "ActionName"}, "Id", "ActionName", "Icon")
	this.Data["Meun"] = list
	this.TplName = "index.html"
}
