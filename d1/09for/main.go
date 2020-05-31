package main

import "fmt"

func main() {
	// 基本格式
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//变种1
	var i = 5
	for ; i < 10; i++ {
		fmt.Println(i)
	}
	//变种2
	var j = 5
	for j < 10 { // 这里不增加
		fmt.Println(j)
		j++ // 如果这里也不写，就是无限循环
	}
	// 无限循环
	/*
		for { 不要运行
			fmt.Println("123")
		}
	*/
	s := "ss撒哈哈"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}
	leftNum := 1
	for ; leftNum < 10; leftNum++ {
		for rightNum := 1; rightNum <= leftNum; rightNum++ {
			fmt.Printf("%d*%d=%d ", leftNum, rightNum, leftNum*rightNum)
		}
		fmt.Printf("\n")
	}
	//跳出一层for，多层for 只跳一层
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // 退出
		}
		fmt.Println(i)
	}
	//跳过一次for循环内部的语句
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue // 5 不打印
		}
		fmt.Println(i)
	}
}
