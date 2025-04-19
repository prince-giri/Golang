package main

import "fmt"

// iota is a predeclared identifier in Go, used only inside const blocks.
// It helps you define incrementing numbers easily, starting from 0.

func main() {
	const (
		Monday = iota + 1  // start from 1 instead of 0
		Tuesday
		Wednesday
		Thursday
		Friday
	)

	fmt.Println(const)
}


// notes:
// Feature	Behavior
// Starts from	0 in each const block
// Increments	Automatically by 1
// Use case	Enum-like constants, bit flags
// Can skip value	Use _
// Works with math	Yes (iota * 10, 1 << iota, etc.)

