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

	// %d ：十进制
	// %b ：二进制
	// %o ：八进制
	// %x ：十六进制
	// %c ：字符
	// %s ：字符串
	// %p ：指针
	// %v ：值
	// %f ：浮点数
	// 通用格式化输出
	// %T ：查看类型
	// %v ： 输出值
	// %#v ：值的 Go 语法表述
	// %+v ：输出类似 %v，输出结构体（struct）时会有字段名
	// %% ：百分号
}
