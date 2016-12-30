package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"strconv"
	"web/utils"
	"crypto/md5"
	"encoding/hex"
)

//用户entity
type User struct {
	Id int				`orm:"pk;auto"`
	Token     string 		`orm:"unique"`
	Username string			`orm:"unique"`
	Password string
	Mobile string			`orm:"null"`
	Status int			`orm:"default(0)"`	//用户状态(0=正常，1=停用)
	Mail string			`orm:"null"`
	Avatar string			`orm:"type(text);default(/static/img/avatar.png)"`
	CreateDate  time.Time 		`orm:"auto_now_add;type(datetime)"`
	ModifyDate  time.Time 		`orm:"auto_now;type(datetime)"`
}

//https://github.com/lyonlai/bootstrap-paginator
func Page(p int, size int) utils.Page{
	o := orm.NewOrm()
	var obj User
	var list []User
	qs := o.QueryTable(obj)
	count, _ := qs.Limit(-1).Count()
	qs.RelatedSel().OrderBy("-CreateDate").Limit(size).Offset((p - 1) * size).All(&list)
	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return utils.PageUtil(c, p, size, list)
}

func Login(username string, password string) (bool, User) {
	o := orm.NewOrm()
	var user User
	h := md5.New()
	h.Write([]byte(password))
	err := o.QueryTable(user).Filter("Username", username).Filter("Password", hex.EncodeToString(h.Sum(nil))).One(&user)
	return err != orm.ErrNoRows, user
}
func FindUserByUserName(username string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Username", username).One(&user)
	return err != orm.ErrNoRows, user
}
func FindUserByToken(token string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Token", token).One(&user)
	return err != orm.ErrNoRows, user
}

func SaveUser(user *User) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(user)
	return id
}

func init() {
	// 需要在init中注册定义的model
	//maxIdle := 30
	//maxConn := 30
	orm.RegisterModel(new(User))
}