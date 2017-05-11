package orm

import (
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
	"fmt"
)

type User struct {
	Id          int
	Name        string
	Profile     *Profile   `orm:"null;rel(one);on_delete(set_null)"` // OneToOne relation
	Post        []*Post `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct {
	Id          int
	Age         int16
	User        *User   `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"`    //设置一对多关系
	Tags  []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:-sAQ=t88@/orm_test?charset=utf8")
	orm.RegisterModel(new(User), new(Profile), new(Tag),new(Post))
}

type UserInfo struct {
	Name string
	Title  string
}
func Orm(){
	o := orm.NewOrm()

	_,err := o.QueryTable("profile").Filter("id",2).Delete()
	if err != nil {
		fmt.Println(err)
	}
}

