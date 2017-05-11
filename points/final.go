package points

import (
	"nas_monitor/common"
	"github.com/goeval"
	"strings"
	"fmt"
	"strconv"
)

var _ = common.Cache

func Final(str1 string) bool{
	//var left,right string

	defer func() {
		if r := recover();r !=nil{
			fmt.Println(r)
			fmt.Println(str1)
		}
	}()
	str := strings.Split(str1,"/")

	if len(str) == 3{
		return false
	}

	if len(str) == 2{
		var val float64
		a1 := str[0][len(str[0])-1:]
		a2 := str[1][:1]

		if a1 != ")" && a2 != "("{

			if len(str[0]) >= 4 {
				_,err := strconv.Atoi(str[0][len(str[0])-4:len(str[0])-3])
				if err == nil {
					a1 = str[0][len(str[0])-4:]
				}else {
					a1 = str[0][len(str[0])-3:]
				}

			}else {
				a1 = str[0][len(str[0])-3:]
			}

			if len(str[1]) >= 4{
				_,err := strconv.Atoi(str[1][3:4])
				if err == nil {
					a2 = str[1][:4]
				}else {
					a2 = str[1][:3]
				}
			}else {
				a2 = str[1][:3]
			}
			val = calculate(fmt.Sprintf("%s/%s",a1,a2))

		}

		if a1 == ")" && a2 != "("{
			if len(str[1]) >= 4{
				_,err := strconv.Atoi(str[1][3:4])
				if err == nil {
					a2 = str[1][:4]
				}else {
					a2 = str[1][:3]
				}
			}else {
				a2 = str[1][:3]
			}
			index := strings.LastIndex(str[0],"(")
			a1 = str[0][index:]
			val = calculate(fmt.Sprintf("%s/%s",a1,a2))

		}

		if a1 != ")" && a2 == "("{
			if len(str[0]) >= 4 {
				_,err := strconv.Atoi(str[0][len(str[0])-4:len(str[0])-3])
				if err == nil {
					a1 = str[0][len(str[0])-4:]
				}else {
					a1 = str[0][len(str[0])-3:]
				}

			}else {
				a1 = str[0][len(str[0])-3:]
			}
			index := strings.Index(str[1],")")
			a2 = str[1][:index+1]
			val = calculate(fmt.Sprintf("%s/%s",a1,a2))

		}

		if a1 ==")" && a2 =="("{
			index := strings.LastIndex(str[0],"(")
			a1 = str[0][index:]
			index = strings.Index(str[1],")")
			a2 = str[1][:index+1]
			val = calculate(fmt.Sprintf("%s/%s",a1,a2))
		}

		value := fmt.Sprintf("%s",val)
		if strings.Contains(value,".") {
			return false
		}
	}

	return true
}

func calculate(expr string) float64{
	s := goeval.NewScope()

	val,err := s.Eval(expr)
	if err != nil {
		fmt.Println(err)
	}
	return val.(float64)
}