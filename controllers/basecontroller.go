package controllers

import (
	"OA/bll"
	"OA/models"
	"OA/utils"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	//实例结构体
	model interface{}
	//实例结构体切片
	list interface{}
	//controller view of file path folder
	pathFolder string
	//result message struct
	message models.ResultMessage
	//bll interface
	iBll    bll.IBll
	admin   bool
	orderBy []string
}

func (this *BaseController) Index() {
	this.TplName = this.pathFolder + "index.html"
}
func (this *BaseController) Add() {
	if this.Ctx.Request.Method == "GET" {
		this.TplName = this.pathFolder + "add.html"
		return
	}
	if error := this.ParseForm(this.model); error != nil {
		this.message.Code = utils.SEND_DATA_ERROR
	} else {
		this.message.Code = this.iBll.Add(this.model)
	}
	this.message.Text = utils.GetCodeText(this.message.Code)
	this.Data["json"] = this.message
	this.ServeJSON()
}

func (this *BaseController) Updata() {
	if this.Ctx.Request.Method == "GET" {
		if error := this.ParseForm(this.model); error != nil {
			this.Data["error"] = "发送的数据有误"
			this.TplName = "error/updatageterror.html"
			return
		}
		if this.iBll.QueryOne(this.model) {
			this.Data["model"] = this.model
			this.TplName = this.pathFolder + "updata.html"
			return
		} else {
			this.Data["error"] = "要修改的数据，已经不存在"
			this.TplName = "error/updatageterror.html"
			return
		}
	}
	if error := this.ParseForm(this.model); error != nil {
		this.message.Code = utils.SEND_DATA_ERROR
	} else {
		this.message.Code = this.iBll.Updata(this.model)
	}
	this.message.Text = utils.GetCodeText(this.message.Code)
	this.Data["json"] = this.message
	this.ServeJSON()
}
func (this *BaseController) Delete() {
	if error := this.ParseForm(this.model); error != nil {
		this.message.Code = utils.SEND_DATA_ERROR
	} else {
		this.message.Code = this.iBll.Delete(this.model)
	}
	this.message.Text = utils.GetCodeText(this.message.Code)
	this.Data["json"] = this.message
	this.ServeJSON()
}
func (this *BaseController) RealDelete() {
	if error := this.ParseForm(this.model); error != nil {
		this.message.Code = utils.SEND_DATA_ERROR
	} else {
		this.message.Code = this.iBll.RealDelete(this.model)
	}
	this.message.Text = utils.GetCodeText(this.message.Code)
	this.Data["json"] = this.message
	this.ServeJSON()
}
func (this *BaseController) GetList() {
	pageSize, _ := this.GetInt("rows")
	pageIndex, _ := this.GetInt("page")
	if pageSize <= 0 {
		pageSize = 10
	}
	if pageIndex <= 0 {
		pageIndex = 1
	}
	where := make(map[string]interface{})
	if !this.admin {
		where["DelFlage"] = false
	}
	total := this.iBll.QueryTable(this.model, this.list, where, []string{}, this.orderBy, pageSize, pageIndex)
	if total > 0 {
		this.message.Code = utils.OPERATE_OK
		this.message.Data = struct {
			row  int64
			data interface{}
		}{row: total, data: this.list}
	} else {
		this.message.Code = utils.QUERY_FAIL
	}
	this.message.Text = utils.GetCodeText(this.message.Code)
	this.Data["json"] = this.message
	this.ServeJSON()
}
