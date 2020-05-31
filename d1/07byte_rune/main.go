package main

import (
	"fmt"
	"unicode"
)

// rune 类型，
// go语言中为了处理非ASCII码类型的字符串，定义了新的rune类型
func main() {
	s := "helloどど" //hello len = 4，
	//len() 求得是byte 字节的数量
	n := len(s)
	fmt.Println(n)

	for i, c := range s {
		fmt.Printf("%v, %c\n", i, c)
	}
	//字符串修改，不能直接修改，需要转换成其他类型
	s2 := "bai"             // 字符串
	s3 := []rune(s2)        // 把字符串强制转换成rune 切片，切片里保存的是字符'b' 'a' 'i'
	s3[0] = 'c'             //字符
	fmt.Println(string(s3)) // 把rune切片强制转换成字符串
	c1 := "a"
	c2 := 'a'
	fmt.Printf("c1:%T, c2:%T, c2的int32的值:%d\n", c1, c2, c2)
	c3 := 'H'
	c4 := byte('H') // byte 本质上是 uint8
	fmt.Printf("c3:%T, c4:%T, c4的uint8的值:%d\n", c3, c4, c4)
	//类型转换
	n1 := 10
	var f float64
	f = float64(n1)
	fmt.Printf("n1: %T, f:%T, f's value: %v", n1, f, f)
	z2 := "hello沙河小王子"
	size := 0
	for _, c := range z2 {
		if unicode.Is(unicode.Han, c) {
			size++
		}
	}
	fmt.Printf("z2中有%d个中文", size)
}
