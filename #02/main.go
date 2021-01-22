package main

import "fmt"

var l = [...]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 6, 6, 7, 7, 8, 8, 9, 9}

func main() {
	res := 0
	for _, n := range l {
		res ^= n
	}
	if res != 0 {
		fmt.Print(res)
	}
}
