package main

import (
	"fmt"
	e "errors"
)
func main(){
   myErr:= e.New("This is an error") //New function must return a pointer to the errorString struct invoked with the string passed to it
 //The errors package has errorString struct which implements the error interface.  
   fmt.Printf("Type of the error is %T\n",myErr) // o/p: *errors.errorString 
   fmt.Printf("value of the error is %v\n",myErr)   
   
}

