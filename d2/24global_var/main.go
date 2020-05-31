package main

import "fmt"

var x = 100 // 定义一个全局变量

func f1() {
	// 变量查找顺序
	// 先在函数内部查找
	// 找不到就往函数外面查找， 一直找到全局
	fmt.Println(x)
}

func main() {
	f1()
}
