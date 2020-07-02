package bll

import (
	"OA/models"
	"OA/utils"
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
		return utils.ADD_ISEXIT_FAIL
	}
	if _, error := this.db.Add(model); error == nil {
		return utils.OPERATE_OK
	}
	return utils.ADD_FAIL
}

// Delete(model interface{}) int
func (this *BaseBll) Delete(model interface{}) int {
	if _, error := this.db.Delete(model); error == nil {
		return utils.OPERATE_OK
	}
	return utils.DEL_FAIL
}

// Updata(model interface{}, cols ...string) int
func (this *BaseBll) Updata(model interface{}, cols ...string) int {
	return 1
}

// RealDelete(model interface{}) int
func (this *BaseBll) RealDelete(model interface{}) int {
	return 1
}

// QueryOne(model interface{}, cols ...string) bool
func (this *BaseBll) QueryOne(model interface{}, cols ...string) bool {
	return true
}

// QueryTable(model, list interface{}, where map[string]interface{}, relatedSel []string, orderBy []string, pageSize, pageIndex int, cols ...string) int64
func (this *BaseBll) QueryTable(model, list interface{}, where map[string]interface{}, relatedSel []string, orderBy []string, pageSize, pageIndex int, cols ...string) int64 {
	return 1
}

// QueryAll(model, list interface{}, where map[string]interface{}, relatedSel []string, orderBy []string, cols ...string)
func (this *BaseBll) QueryAll(model, list interface{}, where map[string]interface{}, relatedSel []string, orderBy []string, cols ...string) {

}
