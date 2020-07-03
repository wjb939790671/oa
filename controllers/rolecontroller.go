package controllers

import (
	"OA/bll"
	"OA/models"
)

type RoleController struct {
	BaseController
}

func (this *RoleController) Prepare() {
	this.model = &models.Role{}
	var list []models.Role
	this.list = &list
	this.iBll = &bll.RoleBll{}
	this.pathFolder = "role/"
	this.unique = "RoleName"
	this.orderBy = []string{"Id", "RoleName"}
}
