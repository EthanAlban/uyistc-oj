package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"unioj/conf"
	"unioj/logs"
	"unioj/models/redisOP"
)

type User struct {
	UId       int    `orm:"column(uid);pk"`
	UserName  string `orm:"column(username);size(60)"`
	UserType  int    `orm:"column(usertype)"`
	RealName  string `orm:"column(realname);size(60)"`
	Password  string `orm:"column(password);size(255)"`
	Credit    int    `orm:"column(credit);"`
	Email     string `orm:"column(email);size(60)"`
	Tel       string `orm:"column(tel);size(20)"`
	Signature string `orm:"column(signature);size(255)"`
	Major     string `orm:"column(major);size(30)"`
	Gitaddr   string `orm:"column(gitaddr);size(60)"`
	Blogaddr  string `orm:"column(blogaddr);size(60)"`
	Avatar    string `orm:"column(avatar);size(60)"`
}

var O orm.Ormer

func init() {
	orm.Debug = true
	dbuser := conf.GetStringConfig("dbuser")
	dbpassword := conf.GetStringConfig("dbpassword")
	dbhost := conf.GetStringConfig("dbhost")
	dblink := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ")/unioj?charset=utf8"
	err := orm.RegisterDataBase("default", "mysql", dblink)
	if err != nil {
		fmt.Printf("err:%v", err)
		logs.LogError("err:mysql connnection failed")
		return
	}
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//beego 注册模型
	orm.RegisterModel(new(User), new(Annnocement), new(ProblemType), new(Problems), new(Tags), new(ProblemTags), new(SysInfo))
	O = orm.NewOrm()
}

// TableName 获取对应数据库表名.
func (u *User) TableName() string {
	return "user"
}

// TableEngine 获取数据使用的引擎.
func (u *User) TableEngine() string {
	return "INNODB"
}

func NewUser() *User {
	return &User{}
}

// GetuserbyUsernamePwd 按手机号和密码查询用户
func (u *User) GetuserbyUsernamePwd(email_or_tel, pwd, usernameType string) (*User, error) {
	user := &User{}
	querySeter := O.QueryTable("user")
	var err error
	if usernameType == "email" {
		err = querySeter.Filter("email", email_or_tel).One(user)
	} else if usernameType == "tel" {
		err = querySeter.Filter("tel", email_or_tel).One(user)
	} else {
		err = querySeter.Filter("username", email_or_tel).One(user)
	}
	if err != nil {
		return nil, err
	}
	if pwd != user.Password {
		return nil, errors.New("wrong password")
	}
	return user, nil
}

// GetUserByUid 根据user id得到user的基本信息  过滤掉系统中的敏感信息
func (u *User) GetUserByUid(uid int) (*User, error) {
	var user User
	if uid < 0 {
		return nil, errors.New("用户id非法")
	}
	res, err := redisOP.RedisGetKey("user_" + strconv.Itoa(uid))
	if err != nil {
		O.Raw("SELECT uid,avatar,username,credit,signature,usertype,gitaddr,major,blogaddr from `user` where uid = ?", uid).QueryRow(&user)
		//O.QueryTable("user").Filter("uid", uid).One(&user)
		str, _ := json.Marshal(user)
		err = redisOP.RedisSetKey("user_"+strconv.Itoa(uid), string(str))
	} else {
		err = json.Unmarshal([]byte(res), &user)
	}
	if err != nil {
		return nil, err
	}
	return &user, err
}

// IsEmailUsed 检查用户输入的email是不是已经被注册过了
func (u *User) IsEmailUsed(email string) bool {
	var user User
	err := O.QueryTable("user").Filter("email", email).One(&user)
	if err != nil {
		if err.Error() == "<QuerySeter> no row found" {
			return false
		}
	}
	return true
}

// IsUserNameUsed  检查用户名是不是被用过了
func (u *User) IsUserNameUsed(username string) bool {
	var user User
	err := O.QueryTable("user").Filter("username", username).One(&user)
	if err != nil {
		if err.Error() == "<QuerySeter> no row found" {
			return false
		}
	}
	return true
}
