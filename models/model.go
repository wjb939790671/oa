package models

import "time"

type ResultMessage struct {
	Code int
	Text string
	Data interface{}
}

type Role struct {
	Id        int
	RoleName  string
	Remark    string
	ParentId  int
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
	IsMeun     bool
	Mothed     bool
	Delflage   bool
	Remark     string
	Roles      []*Role `orm:"reverse(many)"`
}

type Employee struct {
	Id       int
	Name     string
	IdCard   string //id card
	Pwd      string //passward
	Age      int
	Sex      bool //true :男，false：女
	Brithday time.Time

	Roles []*Role `orm:"rel(m2m)"`
}
