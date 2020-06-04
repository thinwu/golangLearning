package main

import "fmt"

// 给自定义类型定义方法
// 只能给自己包里的类型定义方法
type myInt int

func (m myInt) hello() {
	fmt.Println("my int")
}

func main() {
	m := myInt(100)
	m.hello()
}
