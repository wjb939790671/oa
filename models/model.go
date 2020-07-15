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
	Id     int
	Age    int
	Height int
	Weight int

	Name     string
	IdCard   string //id card
	Pwd      string //passward
	Remark   string
	Brithday time.Time
	Delflage bool
	Sex      bool //true :男，false：女

	Roles []*Role `orm:"rel(m2m)"`
}
