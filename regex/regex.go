package regex

import (
	"regexp"
	"fmt"
)


func Regex(){
	p,err := regexp.Compile(`^10\.10\.([0-9]+)\.76$`)
	if err != nil {
		fmt.Println(err)
	}
	b := p.MatchString("10.10.100.76")
	fmt.Println(b)
}

