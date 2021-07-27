package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int
	Username string
	Password string
	Sex      int
	Role     *UserRole `orm:"rel(fk)"`
}

type UserRole struct {
	Id   int
	Name string  `orm:"column(role_name)"`
	Desc string  `orm:"column(role_desc)"`
	User []*User `orm:"reverse(many)"`
}

func init() {
	orm.Debug = true
	orm.RegisterDataBase("default", "mysql", "root:Abc741258963.0@tcp(192.168.112.10:3306)/demo_01?charset=utf8")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//beego 注册模型
	orm.RegisterModel(new(User), new(UserRole))
}

func FIndUsersByUsername(Username string) (*User, error) {
	o := orm.NewOrm() // 创建新的orm
	user := &User{}
	err := o.QueryTable("User").Filter("username", Username).One(user)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到该用户")
	}
	return user, nil
}

func AllUsers() ([]*User, error) {
	var users []*User
	_, err := orm.NewOrm().QueryTable("User").RelatedSel().All(&users)
	if err != nil {
		fmt.Println("错误：err:", err)
		return nil, err
	}
	return users, nil
}
