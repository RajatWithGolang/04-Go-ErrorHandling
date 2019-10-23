package main

import "fmt"

func defFooStart() { 
	fmt.Println("defFooStart() executed") //step 6
}

func defFooEnd() {
	fmt.Println("defFooEnd() executed")
}

func defMain() {
	fmt.Println("defMain() executed")
}

func foo() {
	fmt.Println("foo() executed") // step 3
	
	defer defFooStart() // step 5
	
	panic("panic from foo()")  // step 4
	
	defer defFooEnd() 
	
	fmt.Println("foo() done")
}

func main() {
	fmt.Println("main() started")  // step 1
	
	defer defMain() // step 7
	
	foo()  //step 2 
	
	fmt.Println("main() done")
}
