package main


//表达式
//author:Xiong Chuan Liang
//date:2015-2-2

import (
	"fmt"
	"github.com/xclpkg/algorithm"
	"strconv"
	"errors"
)

func main(){

	// 中序表达式       后序表达式
	// a+b            = ab+
	// (a+b)/c        = ab+c/
	// a+b*(c+d)      = abcd+*+
	// a*b+c*(d-e)/f  = ab*cde-*f/+

	//str := "a*b+c*(d-e)/f"
	//str := "1*2+3*(4-5)/6"

	str := "1*(1*6+13)"

	exp,err := ExpConvert(str)
	fmt.Println(exp)
	if err != nil {
		fmt.Println("中序表达式转后序表达式失败! ",err)
	}else{
		Exp(exp)
	}

	//v := 1*2+3*(4-5)/6
	v := 1*(1*6+13)
	fmt.Println("标准结果: ",v)

}


func ExpConvert(str string)(string,error){

	var result string
	stack := algorithm.NewStack()
	for _,s := range str {
		ch := string(s)
		if IsOperator(ch) { //是运算符

			if stack.Empty() || ch == "(" {
				stack.Push(ch)
			}else{
				if ch == ")" { //处理括号
					for{
						if stack.Empty() {
							return "",errors.New("表达式有问题! 没有找到对应的\"(\"号")
						}
						if stack.Top().(string) == "(" {
							break
						}
						result += stack.Top().(string)
						stack.Pop()
					}

					//弹出"("
					stack.Top()
					stack.Pop()
				}else{ //非括号
					//比较优先级
					for{
						if stack.Empty() {
							break
						}
						m := stack.Top().(string)
						if Priority(ch) > Priority(m) {
							break
						}
						result += m
						stack.Pop()
					}
					stack.Push(ch)
				}
			}

		}else{	//非运算符
			result += ch
		} //end IsOperator()

	} //end for range str

	for {
		if stack.Empty() {
			break
		}
		result += stack.Top().(string)
		stack.Pop()
	}

	//fmt.Println("ExpConvert() str    = ",str )
	//fmt.Println("ExpConvert() result = ",result)
	return result,nil
}

func Exp(str string){
	//fmt.Println("\nCalc \nExp() :",str )
	stack := algorithm.NewStack()
	for _,s := range str {
		ch := string(s)
		if IsOperator(ch) { //是运算符
			if stack.Empty() {
				break
			}
			b := stack.Top().(string)
			stack.Pop()

			a := stack.Top().(string)
			stack.Pop()

			ia,ib := convToInt32(a,b)
			sv := fmt.Sprintf("%d",Calc(ch,ia,ib))
			stack.Push( sv )
			//fmt.Println("Exp() ",a,"",ch,"",b,"=",sv)
		}else{
			stack.Push(ch)
			//fmt.Println("Exp() ch: ",ch)
		} //end IsOperator
	}

	//stack.Print()
	if !stack.Empty() {
		fmt.Println("表达式运算结果: ", stack.Top() )
		stack.Pop()
	}
}

func convToInt32(a,b string)(int32,int32){

	ia, erra := strconv.ParseInt(a, 10, 32)
	if erra != nil {
		panic(erra)
	}

	ib, errb := strconv.ParseInt(b, 10, 32)
	if errb != nil {
		panic(errb)
	}
	return int32(ia),int32(ib)
}

func IsOperator(op string)(bool){
	switch(op){
	case "(",")","+","-","*","/":
		return true
	default:
		return false
	}
}

func Priority(op string)(int){
	switch(op){
	case "*","/":
		return 3
	case "+","-":
		return 2
	case "(":
		return 1
	default:
		return 0
	}
}

func Calc(op string,a,b int32)(int32){

	switch(op){
	case "*":
		return (a * b)
	case "/":
		return (a / b)
	case "+":
		return (a + b)
	case "-":
		return (a - b)
	default:
		return 0
	}
}