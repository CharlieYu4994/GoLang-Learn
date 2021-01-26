package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().UnixNano()
	now := time.Now().Format("2006/01/02 15:04:05")
	fmt.Println(now)
	end := time.Now().UnixNano()
	last := (end - start) / 1000
	fmt.Printf("lasts: %dus\n", last)
}
