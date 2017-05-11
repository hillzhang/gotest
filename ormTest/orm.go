package ormTest

import (
	"sys_monitor/astaxie/beego/orm"
	_"sys_monitor/go-sql-driver/mysql"
	"fmt"
	"time"
	"sort"
)

var table = `
CREATE TABLE IF NOT EXISTS store_%s (
 id INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
 ts INT(11) NOT NULL COMMENT '时间戳',
 value INT(5) UNSIGNED NOT NULL COMMENT '访问数',
 PRIMARY KEY (id))
 ENGINE = InnoDB
 AUTO_INCREMENT = 1
 DEFAULT CHARACTER SET = utf8
 `

type StoreInfo struct {
	Id int
	Ts int64
	Value int32
}

func TableName() string {
	return "store_"+time.Now().Format("200601")
}

func init() {
	orm.RegisterDataBase("default","mysql","root:-sAQ=t88@tcp(localhost:3306)/varnish?charset=utf8")
	orm.RegisterModel(new(StoreInfo))
}
var(
	sql_insert = "insert into %s(ts,value) values(?,?)"
	sql_select = "select * from %s"
)


func Orm(){
	o := orm.NewOrm()

	//_,err := o.Raw(fmt.Sprintf(table,time.Now().Format("200601"))).Exec()
	//if err != nil {
	//	fmt.Println("raw",err)
	//}
	//result,err := o.Raw(fmt.Sprintf(sql_insert,TableName()),time.Now().Unix(),7).Exec()
	//if err != nil {
	//	fmt.Println(err)
	//}else {
	//	fmt.Println(result)
	//}
	//info := new(StoreInfo)
	//info.Ts = time.Now().Unix()
	//info.Value = 6
	//
	//_,err = o.Insert(info)
	//if err != nil {
	//	fmt.Println("insert",err)
	//}

	users := []StoreInfo{}

	count,err := o.Raw(fmt.Sprintf(sql_select,TableName())).QueryRows(&users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(count,users)


}

