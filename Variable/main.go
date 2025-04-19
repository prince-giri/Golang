package main

import (
	"fmt"
)

func main() {
	// Variable declrearation
	var first_name string
	var last_name string
	var age uint64

	// variable value assigned
	first_name = "Prince"
	last_name = "Giri"
	age = 26

	var full_name string = first_name + " " + last_name

	fmt.Printf("My name is: %v \nMy current age is: %v", full_name, age)
}

Variable Important Notes:

// 1. Uninitialized variables are given its zero value (e.g int = 0, string = "", bool = false)

// 2. Go Primitive Types
// The possible values for bool are true and false.

// uint8 : unsigned 8-bit integers (0 to 255)
// uint16 : unsigned 16-bit integers (0 to 65535)
// uint32 : unsigned 32-bit integers (0 to 4294967295)
// uint64 : unsigned 64-bit integers (0 to 18446744073709551615)
// int8 : signed 8-bit integers (-128 to 127)
// int16 : signed 16-bit integers (-32768 to 32767)
// int32 : signed 32-bit integers (-2147483648 to 2147483647)
// int64 : signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
// int is either int64 or int32 depends on the implementation.
// float32 : set of all IEEE-754 32-bit floating-point numbers
// float64 : set of all IEEE-754 64-bit floating-point numbers
// complex64 the set of all complex numbers with float32 real and imaginary parts
// complex128 the set of all complex numbers with float64 real and imaginary parts
// byte alias for uint8 rune alias for int32

// 3. Visibility
// If variables/functions starts with Uppercase character, it is accessible outside the scope of its package, if lowercase then it is only accessible inside its package.

// package myPkg

// var Uppercase = "This is accessible outside the pkg"
// var lowercase = "This is not accessible outside the pkg"
// func UppercaseFunc() string  {	return "This is accessible outside the pkg"}
// func lowercaseFunc() string  {	return "This is accessible outside the pkg"}

// // Another file:
// package main
// import "myPkg"

// func main() {
//   //Accessible
//   println(myPkg.Uppercase)
//   println(myPkg.UppercaseFunc())

//   //Not Accessible
//   println(myPkg.lowercase)
//   println(myPkg.lowercaseFunc())
// }
