package main

import "fmt"

func main() {
	var i1 = 12            // 类型推导
	fmt.Printf("%d\n", i1) // 10进制输出
	fmt.Printf("%b\n", i1) // 二进制输出
	fmt.Printf("%o\n", i1) // 8进制输出
	fmt.Printf("%x\n", i1) // 16进制输出

	//八进制
	i2 := 077
	fmt.Printf("%d\n", i2)
	//十六进制
	i3 := 0xf
	fmt.Printf("%d\n", i3)
	//查看类型
	fmt.Printf("%T\n", i3)
	// 查看字符
	fmt.Printf("%c\n", 'a')
	//查看字符串
	fmt.Printf("%s\n", "aa")
	//声明int8 类型的变量
	i4 := int8(9) //明确指定类型，否则就是默认
	fmt.Printf("%T", i4)
}
