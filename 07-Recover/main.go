package main

import "fmt"

func defFooStart() {
	fmt.Println("defFooStart() started")  //step 6
	r := recover()
	if r!=nil{
		fmt.Println("program with panicking with",r) //step 7
	}
	fmt.Println("defFooStart() executed")  //step 8
}

func defFooEnd() {
	fmt.Println("defFooEnd() executed")
}

func defMain() {
	fmt.Println("defMain() executed")  // step 12
}

func foo() {
	fmt.Println("foo() executed")  //step 3
	
	defer defFooStart() // step 5
	
	panic("foo() panics")  //step 4
	
	defer defFooEnd()  //step 9
	
	fmt.Println("foo() done")
}

func main() {
	fmt.Println("main() started") // Step 1
	
	defer defMain() //step 11
	
	foo() //step 2
	
	fmt.Println("main() done") //step 10
}
