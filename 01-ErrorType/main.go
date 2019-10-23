package main

import (
	"fmt"
)
type myError string

func (me *myError) Error() string{
	return "Here is some error"
}
func main(){
  var me1 myError
  fmt.Println(me1.Error())
}
