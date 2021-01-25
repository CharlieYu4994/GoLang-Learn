package main

import "fmt"

var l = [...]int{1, 3, 5, 7, 8}

func main() {
	s := 0
	for _, v := range l {
		s += v
	}
	fmt.Println(s)

	s = 0
	for i, v1 := range l {
		for j, v2 := range l {
			if i < j {
				if v1+v2 == 8 {
					fmt.Printf("(%d, %d)\t", i, j)
				}
			}
		}
	}
}
