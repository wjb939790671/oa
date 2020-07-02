package utils

const (
	ADD_OK    = 10000
	DEL_OK    = 10001
	UPDATA_OK = 10002
	QUERY_OK  = 10003

	ADD_ISEXIST      = 40000
	ADD_FAIL         = 40001
	DEL_FAIL         = 40002
	DEL_NOT_EXIST    = 40003
	UPDATA_FAIL      = 40004
	UPDATA_NOT_EXIST = 40005
	UPDATA_UNIQUE    = 40006
	SEND_DATA_ERROR  = 40007
	QUERY_FAIL       = 40008
)

func GetCodeText(code int) string {
	return map[int]string{
		ADD_OK:           "添加成功",
		ADD_ISEXIST:      "添加失败，数据已存在",
		ADD_FAIL:         "添加失败",
		UPDATA_OK:        "修改成功",
		UPDATA_FAIL:      "修改失败",
		UPDATA_NOT_EXIST: "修改失败，数据不存在",
		SEND_DATA_ERROR:  "数据有误，请核实要操作的数据",
		QUERY_FAIL:       "查询失败",
		DEL_OK:           "删除成功",
		DEL_FAIL:         "删除失败",
		DEL_NOT_EXIST:    "删除失败，数据不存在",
		UPDATA_UNIQUE:    "数据已存在",
		QUERY_OK:         "查询成功",
	}[code]
}
