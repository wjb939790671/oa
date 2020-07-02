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
func (this *MainController) LoadMeun() {
	meunId, error := this.GetInt("meunId")
	if error != nil {
		return
	}
	tree := getTree(this.model, this.iBll, meunId)
	this.Data["json"] = tree
	this.ServeJSON()

}

type Tree struct {
	Id         int                    `json:"id"`
	Text       string                 `json:"text"`
	IconCls    string                 `json:"iconCls"`
	State      string                 `json:"State"`
	Checked    bool                   `json:"checked"`
	Attributes map[string]interface{} `json:"attributes"`
	Children   []Tree                 `json:"children"`
}

func getTree(model interface{}, iBll bll.IBll, parentId int) []Tree {
	where := make(map[string]interface{})
	where["ParentId"] = parentId
	var list []models.Action
	iBll.QueryAll(model, &list, where, []string{"Id", "ActionName"}, "Id", "ActionName", "Icon", "Url", "IsMeun", "M")
	var trees []Tree
	for _, value := range list {
		var tree Tree
		if value.IsMeun {
			tree.Children = getTree(model, iBll, value.Id)
		}
		atrr := make(map[string]interface{})
		atrr["url"] = value.Url
		tree.Id = value.Id
		tree.Text = value.ActionName
		tree.IconCls = value.Icon
		trees = append(trees, tree)
	}
	return trees
}
