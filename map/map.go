package main

import (
	"fmt"
	"maps"
)

func main() {
	fmt.Println("Started to learn Map!")
	m := make(map[string]int)
	fmt.Println("Initalized hash is", m)
	fmt.Printf("Type of m is %T", m)
	m["a"] = 0
	m["b"] = 2
	m["c"] = 3
	fmt.Println("Initalized hash is", m)

	for k, v := range m {
		fmt.Println(k, v)
		if k == "b" {
			m[k] = 10
		}

	}
	fmt.Println("Initalized hash is", m)
	delete(m, "b")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 3, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
