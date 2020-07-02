package main

import (
	"OA/models"
	_ "OA/routers"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	beego.Run()
}
func init() {
	// var result Result

	// getXml(&result.Root)
	// CreateFile(&result.Root.Tokens[0])

	// 参数1        数据库的别名，用来在 ORM 中切换数据库使用
	// 参数2        driverName
	// 参数3        对应的链接字符串

	//orm.RegisterDataBase("default", "mysql", "root:123@/oa?charset=utf8")

	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)
	maxIdle := 2
	maxConn := 3
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysql"), maxIdle, maxConn)
	//orm.RegisterDataBase("default", "mysql", "username:password@tcp(127.0.0.1:3306)/db_name?charset=utf8", 30)

	// 需要在init中注册定义的model
	orm.RegisterModel(new(models.Employee), new(models.Role), new(models.Action))

	// create table
	orm.RunSyncdb("default", false, true)

}

func getXml(token *Token) {
	file, error := os.Open("models/customTable.xml")
	if error != nil {
		fmt.Println("read file fail :", error)
		return
	}
	decoder := xml.NewDecoder(file)
	parseToken(decoder, token)
}

type Attribute struct {
	Name  string
	Value string
}
type Token struct {
	Name       string
	Text       string
	Attributes []Attribute
	Tokens     []Token
}
type Result struct {
	Root Token
}

func parseToken(decoder *xml.Decoder, myToken *Token) {
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("parse file fail:", err)
			return
		}
		switch tp := token.(type) {
		//start element
		case xml.StartElement:
			se := xml.StartElement(tp)
			var newToken Token
			newToken.Name = se.Name.Local
			for _, attr := range se.Attr {
				newToken.Attributes = append(newToken.Attributes, Attribute{Name: attr.Name.Local, Value: attr.Value})
			}
			parseToken(decoder, &newToken)
			myToken.Tokens = append(myToken.Tokens, newToken)
			//end element
		case xml.EndElement:
			return
			//content
		case xml.CharData:
			cd := xml.CharData(tp)
			data := strings.TrimSpace(string(cd))
			myToken.Text = data
		}
	}
}

func CreateFile(token *Token) {
	fmt.Println(token)
	strs := ""
	if token.Name != "model" {
		return
	}
	for _, value := range token.Tokens {
		str := "type "
		str += value.Name + " struct{\r\n"
		str += "Base     `orm:\" - \"`\r\n"
		if len(value.Tokens) > 0 {
			for _, v := range value.Tokens {
				str += v.Name + " "
				if len(v.Attributes) > 0 {
					str += v.Attributes[0].Value + " \r\n"
				} else {
					str += " \r\n"
				}
			}
		}
		str += "}\r\n"
		strs += str

	}
	fmt.Println(strs)
	file, err := os.OpenFile("models/automodel.go", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		fmt.Println("atuo model file open fail:", err)
		return
	}
	defer file.Close()
	file.WriteString(strs)
}
