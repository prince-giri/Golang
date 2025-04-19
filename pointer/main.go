//1. A pointer is a variable that stores the memory address of another variable.
// Instead of holding a value like 10, it holds where in memory that value is stored.

//2.  A pointer declared but not initialized has a zero value of nil.

package main

import "fmt"

func main() {
	a := 10
	p := &a

	fmt.Printf("Value of a is: %v \n", a)
	fmt.Printf("Type of a is %T \n", a)
	fmt.Println(p)
	fmt.Printf("Type of a is %T \n", p)
	b := *p
	fmt.Printf("Value of b is: %v \n", b)
}

// Pointers Notes
// var value int = 1000
// var pointer *int = &value
// println(value)                //1000
// println(pointer)              //0xfffffffff
// println(*pointer)             //1000
// (*pointer)++		  			      //1001
// *pointer = *pointer + 10	    //1011
// println(*pointer)			        //1011
// println(*pointer + *pointer)  //1011 + 1011 = 2022
