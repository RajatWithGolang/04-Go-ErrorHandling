package main

import (
	"fmt"
	"errors"
)
func divideByZero(a int,b int) (int,error){
	if b == 0 {
		return 0, errors.New("Can not devide by Zero!")
	} else {
		return (a / b), nil
	}
}
func main(){
	result,err := divideByZero(4,0)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(result)
}