package utils

const (
	OPERATE_OK            = 4000
	ADD_ISEXIT_FAIL       = 5000
	ADD_FAIL              = 5001
	DEL_FAIL              = 5002
	DEL_NOTISEXIT_FAIL    = 5003
	UPDATA_FAIL           = 5004
	UPDATA_NOTISEXIT_FAIL = 5005
	SEND_DATA_ERROR       = 5006
	QUERY_FAIL            = 5007
	OPERATE_FAIL          = 5008
)

func GetCodeText(code int) string {
	return map[int]string{
		OPERATE_OK:            "操作成功",
		ADD_ISEXIT_FAIL:       "添加失败，数据已存在",
		ADD_FAIL:              "添加失败",
		UPDATA_FAIL:           "修改失败",
		UPDATA_NOTISEXIT_FAIL: "修改失败，数据不存在",
		SEND_DATA_ERROR:       "数据有误，请核实要操作的数据",
		QUERY_FAIL:            "查询失败",
		OPERATE_FAIL:          "操作失败",
	}[code]
}
