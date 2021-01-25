package main

import "fmt"

func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if i > j {
				continue
			}
			fmt.Printf("%dx%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}
}
