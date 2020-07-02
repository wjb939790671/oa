package bll

import (
	"OA/models"
	"OA/utils"

	"github.com/astaxie/beego/orm"
)

type ActionBll struct {
	BaseBll
}

type BaseBll struct {
	db models.DataBase
}

// Add(model interface{}) int
func (this *BaseBll) Add(model interface{}, uniqueCols ...string) int {
	if this.db.IsExist(model, uniqueCols...) {
		return utils.ADD_ISEXIST
	}
	if _, error := this.db.Add(model); error == nil {
		return utils.ADD_OK
	}
	return utils.ADD_FAIL
}

// Delete(model interface{}) int
func (this *BaseBll) Delete(model interface{}) int {
	_, error := this.db.Delete(model)
	if error == nil {
		return utils.DEL_OK
	}
	if error == orm.ErrNoRows {
		return utils.DEL_NOT_EXIST
	}
	return utils.DEL_FAIL
}

// Updata(model interface{}, cols ...string) int
func (this *BaseBll) Updata(model interface{}, cols ...string) int {
	_, error := this.db.Updata(model, cols...)
	if error == nil {
		return utils.UPDATA_OK
	}
	if error == orm.ErrNoRows {
		return utils.UPDATA_NOT_EXIST
	}
	return utils.UPDATA_FAIL
}

// RealDelete(model interface{}) int
func (this *BaseBll) RealDelete(model interface{}) int {
	_, error := this.db.RealDelete(model)
	if error == nil {
		return utils.DEL_OK
	}
	if error == orm.ErrNoRows {
		return utils.DEL_NOT_EXIST
	}
	return utils.DEL_FAIL
}

// QueryOne(model interface{}, cols ...string) bool
func (this *BaseBll) QueryOne(model interface{}, cols ...string) bool {
	error := this.db.QueryOne(model, cols...)
	if error == nil {
		return true
	}
	return false
}

// QueryTable(model, list interface{}, where map[string]interface{}, relatedSel []string, orderBy []string, pageSize, pageIndex int, cols ...string) int64
func (this *BaseBll) QueryTable(model, list interface{}, where map[string]interface{}, relatedSel []interface{}, orderBy []string, pageSize, pageIndex int, cols ...string) int64 {
	total, _ := this.db.QueryTable(model, list, where, relatedSel, orderBy, pageSize, pageIndex, cols...)
	return total
}

// QueryAll(model, list interface{}, where map[string]interface{},  orderBy []string, cols ...string)
func (this *BaseBll) QueryAll(model, list interface{}, where map[string]interface{}, orderBy []string, cols ...string) {
	this.db.QueryAll(model, list, where, orderBy, cols...)
}

