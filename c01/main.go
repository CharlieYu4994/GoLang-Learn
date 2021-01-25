package main

import "fmt"

var a = 1
var b = 2.33
var c = true
var d = "Hello World"
var e = "这里是 Go 语言学习使用的代码"

func main() {
	fmt.Printf("%T %v \n", a, a)
	fmt.Printf("%T %v \n", b, b)
	fmt.Printf("%T %v \n", c, c)
	fmt.Printf("%T %v \n", d, d)

	s := 0
	for _, v := range e {
		if v > 126 {
			s++
		}
	}
	fmt.Println(s)
}
