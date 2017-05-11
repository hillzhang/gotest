package sort

import (
	"sort"
	"fmt"
	"time"
)

type User struct {
	Name string
	Age int
}

type Users []*User

func (u Users) Len()int{
	return len(u)
}

func (u Users) Less(i,j int)bool{
	if u[i].Age < u[j].Age{
		return true
	}else if u[i].Age > u[j].Age {
		return false
	}else {
		return u[i].Name < u[j].Name
	}
}

func (u Users) Swap(i,j int){
	var temp *User = u[i]
	u[i] = u[j]
	u[j] = temp
}

func Do(){
	u1 := &User{"hillzhang",25}
	u2 := &User{"xillzhang",27}
	u3 := &User{"zillzhang",22}
	u4 := &User{"gillzhang",22}

	plist := Users([]*User{u1,u2,u3,u4})
	sozhirt.Sort(plist)
	for _,val := range plist{
		fmt.Println(*val)
	}
	now := time.Now()
	next := time.Date(now.Year(),now.Month(),now.Day(),0,0,0,0,time.Local)
	fmt.Println(next.Format("2006-01-02 15:04:05"))
}