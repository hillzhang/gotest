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
	//orm.RunSyncdb("default",false,true)
}

type UserInfo struct {
	Name string
	Title  string
}
func Orm(){
	o := orm.NewOrm()
	//_,err := o.QueryTable("profile").Filter("id",2).Delete()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//profile := new(Profile)
	//profile.Age = 30
	//
	//user := new(User)
	//user.Profile = profile
	//user.Name = "slene"
	//user := User{Name: "slene"}
	//if created, id, err := o.ReadOrCreate(&user,"Name"); err == nil {
	//	if created {
	//		fmt.Println("New Insert an object. Id:", id)
	//	} else {
	//		fmt.Println("Get an object. Id:", id)
	//	}
	//}

	//fmt.Println(o.Insert(profile))
	//fmt.Println(o.Insert(user))
	//user := User{Id:1}
	//count ,err := o.InsertMulti(5,[]Post{
	//	{Title:"go",User:&user},
	//	{Title:"python",User:&user},
	//	{Title:"java",User:&user},
	//	{Title:"lua",User:&user},
	//	{Title:"js",User:&user},
	//})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(count)

	//var lists []orm.ParamsList
	//
	//num, err := o.QueryTable("post").Filter("User__id",1).ValuesList(&lists,"Title")
	//if err == nil {
	//	fmt.Printf("Result Nums: %d\n", num)
	//	for _, row := range lists {
	//		fmt.Println(row[0])
	//	}
	//}
	//user := &User{}
	//o.QueryTable("user").Filter("Id", 1).RelatedSel().One(user)
	//// 自动查询到 Profile
	//fmt.Println(user.Profile)
	//// 因为在 Profile 里定义了反向关系的 User，所以 Profile 里的 User 也是自动赋值过的，可以直接取用。
	//fmt.Println(user.Profile.User)
	//var posts []*Post
	//	num, err := o.QueryTable("post").Filter("User", 1).RelatedSel().All(&posts)
	//if err == nil {
	//	fmt.Printf("%d posts read\n", num)
	//	for _, post := range posts {
	//		fmt.Printf("Id: %d, UserName: %s, Title: %s\n", post.Id, post.User.Name, post.Title)
	//	}
	//}


	//user := User{Id: 1}
	//err := o.Read(&user)
	//
	//num, err := o.LoadRelated(&user, "Profile")
	//if err != nil {
	//	fmt.Println(err)
	//}else {
	//	fmt.Println(num)
	//}
	//fmt.Println(user.Profile)

	//var titles []string
	//num,err := o.Raw("select title from post where user_id = ?").SetArgs(1).QueryRows(&titles)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(num,titles)

	var users []User
	qb,_ := orm.NewQueryBuilder("mysql")
	qb.Select("user.name",
	"profile.age").From("user").
	InnerJoin("profile").On("user.profile_id = profile.id").
	Where("profile.age>?").OrderBy("name").Desc().
	Limit(10).Offset(0)

	sql := qb.String()
	i,err :=o.Raw(sql,10).QueryRows(&users)
	if err != nil {
		fmt.Println(err)

	}else {fmt.Println(i)}
	for _,val := range users{
		fmt.Println(val)
	}
}

