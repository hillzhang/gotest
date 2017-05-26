package error

import "fmt"

type MyError struct {
	Msg string
}

const (

)
func NewError(err string) error{
	return &MyError{err}
}

func (this *MyError) Error() string{
	return this.Msg
}
var _ error = new(MyError)

func Err(){
	e := NewError("this is a error test")
	fmt.Println(e.Error())
}
