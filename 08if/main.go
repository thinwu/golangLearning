package main

import "fmt"

func main() {
	age := 10
	if age > 18 {
		fmt.Println("aaa")
	} else {
		fmt.Println("bbb")
	}
	if age2 := 9; age2 > 2 {
		fmt.Println("aaa")
	} else {
		fmt.Println("bbb")
	}
	fmt.Println(age)
	//fmt.Println(age2) 在这里找不到，作用域只存在条件判断里，用完就内存回收，节约空间
}
