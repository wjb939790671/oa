package bll

type IBll interface {
	//@param model 实例结构体的地址
	Add(model interface{}) int
	Delete(model interface{}) int
	Updata(model interface{}, cols ...string) int
	RealDelete(model interface{}) int
	QueryOne(model interface{}, cols ...string) bool
	QueryTable(model, list interface{}, where map[string]interface{}, relatedSel []string, orderBy []string, pageSize, pageIndex int, cols ...string) int64
	QueryAll(model, list interface{}, where map[string]interface{}, orderBy []string, cols ...string)
	//M2MAdd(model, dataModel interface{}, relatedSel string) (int64, error)
	//M2MQuery(model interface{}, relatedSel string) (int64, error)
	// IsExist(model interface{}, cols ...string) bool
}
