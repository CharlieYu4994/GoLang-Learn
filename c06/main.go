package main

import (
	"fmt"
	"strings"
)

var s = "I have a pan, and the pan is red"

func main() {
	s = strings.ReplaceAll(s, ",", "")
	words := strings.Split(s, " ")
	stat := make(map[string]int, 10)
	for _, word := range words {
		_, ok := stat[word]
		if !ok {
			stat[word] = 0
		}
		stat[word]++
	}
	fmt.Printf("%#v", stat)
}
