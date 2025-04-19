package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("What's is you name?:")
	// var name string
	// fmt.Scan(&name)
	// fmt.Println("My name is", name)
	// But the proble with scan is in only read till whitespec

	reader := bufio.NewReader((os.Stdin))
	name, _ := reader.ReadString('\n')
	println("My name is:", name)

}
