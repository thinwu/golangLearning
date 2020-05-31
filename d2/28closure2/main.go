package main

import "fmt"

// 闭包是一个函数，这个函数包含了他外部作用域变量的引用
// 低层原理
// 1. 函数可以作为返回值
// 2. 函数内部查找变量的顺序，先在自己内部找， 找不到往外层找

func adder(x int) func(int) int {
	//var x = 100
	return func(y int) int {
		x += y
		return x
	}
}
func main() {
	ret := adder(100)
	/*
		ret = func(int) int
	*/
	ret2 := ret(200)
	/*
		y = 200
		x = 100
	*/
	fmt.Println(ret2)
}
