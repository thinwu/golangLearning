package main

import "fmt"

// 给自定义类型定义方法
// 只能给自己包里的类型定义方法
type myInt int

func (m myInt) hello() {
	fmt.Println("my int")
}

func main() {
	m := myInt(100) // 类型强制转换
	// 1. var m myInt = 100
	// 2. var m myInt
	// m = 100
	// 3. m := myInt(100)
	// 4. var m = myInt(100)
	m.hello()
}
