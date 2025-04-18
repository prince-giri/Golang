package main

import (
	"fmt"
	"learning/hello"
	"learning/hello/hello2"
)

func main() {
	hello.SayHello()
	fmt.Println("Hello World!")
	hello.Goodbye()
	hello2.Hello2()
}
