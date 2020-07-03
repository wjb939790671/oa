package controllers

import (
	"OA/bll"
	"OA/models"
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
}
