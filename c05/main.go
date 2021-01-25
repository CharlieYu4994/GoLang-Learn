package main

import (
	"fmt"
	"sort"
)

func main() {
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)

	b := [...]int{3, 7, 8, 9, 1}
	c := b[:]
	sort.Ints(c)
	fmt.Println(c)
}
