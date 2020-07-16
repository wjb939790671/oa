package models

import "time"

type ResultMessage struct {
	Code int         `json:"code"`
	Text string      `json:"text"`
	Data interface{} `json:"data"`
}

type Role struct {
	Id        int
	RoleName  string
	Remark    string
	ParentId  int
	Delflage  bool
	Actions   []*Action   `orm:"rel(m2m)"`
	Employees []*Employee `orm:"reverse(many)"`
}
type DataBase struct{}
type Action struct {
	Id         int
	ActionName string
	Url        string
	Icon       string
	ParentId   int //上级菜单0为顶级菜单
	IsMenu     bool
	Mothed     bool
	Delflage   bool
	Remark     string
	Roles      []*Role `orm:"reverse(many)"`
}

type Employee struct {
	Id         int
	Position   int //职位
	Delflage   bool
	Sex        bool   //true :男，false：女
	JobNumber  string //工号
	Name       string
	Pwd        string //passward
	Remark     string
	Phone      string
	EnteryTime time.Time //入职日期

	EmployeeExtend *EmployeeExtend `orm:"rel(one)"`
	Roles          []*Role         `orm:"rel(m2m)"`
}
type EmployeeExtend struct {
	Age                 int
	Height              int
	Weight              int
	IdCard              string //id card
	EducationBackground string //学历
	Nation              string //民簇
	NativePlace         string //籍贯
	PresentAddress      string //现居地址
	WorkExperience      string //工作经历
	Brithday            time.Time
	Employee            *Employee `orm:"reverse(one)`
}
