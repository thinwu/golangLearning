package main

import "fmt"

// fmt 占位符
func main() {
	var n = 100
	s := "hello"
	//查看类型
	fmt.Printf("%T\n", n)
	//查看变量的值
	fmt.Printf("%v\n", n)
	// 查看字符串
	fmt.Printf("字符串: %s\n", s)
	fmt.Printf("字符串: %#v\n", s) //# 会带一个具体描述符，字符串会带引号

}
