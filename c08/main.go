package main

import "fmt"

func main() {
	a := new(int)
	*a = 100
	fmt.Println(*a)

	b := make(map[string]int, 1)
	b["hello"] = 100
	fmt.Println(b)
}
