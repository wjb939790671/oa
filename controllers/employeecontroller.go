package controllers

import (
	"OA/bll"
	"OA/models"
)

type EmployeeController struct {
	BaseController
}

func (this *EmployeeController) Prepare() {
	this.model = &models.Employee{}
	var list models.Employee
	this.list = &list
	this.unique = "IdCard"
	this.iBll = &bll.EmployeeBll{}
}
