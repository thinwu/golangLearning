package main

import "fmt"

func funcA() {
	fmt.Println("a")
}

func funcB() {
	defer func() { // defer 要在可能出现panic 之前
		// recover 要和panic 成对出现
		err := recover() // 注释掉这两句看区别
		fmt.Println(err) // 注释掉这两句看区别
		fmt.Println("释放资源")
	}()
	panic("出现了严重错误")
	fmt.Println("b")
}

func funcC() {
	fmt.Println("c")
}

func main() {
	funcA()
	funcB()
	funcC()
}
