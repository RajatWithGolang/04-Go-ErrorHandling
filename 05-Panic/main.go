package main

import "fmt"
func foo(){
	fmt.Println("Foo Started")
	panic("Exception Occured")
	fmt.Println("Foo Terminated")
}

func main(){
	fmt.Println("Main Started")
	foo()
	fmt.Println("Main Terminated")
}