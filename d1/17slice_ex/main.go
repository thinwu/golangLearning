package main

import "fmt"

func main() {
	var a = make([]string, 5, 10) // 初始化有5个长度，容量10
	fmt.Println(a)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprint(i)) // append 从第六个位置开始
	}
	fmt.Println(a)
}
