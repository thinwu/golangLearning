package main

import (
	"fmt"
)

func main() {
	//math.MaxFloat32 // float32 最大值
	f1 := 1.234556
	fmt.Printf("%T\n", f1) //默认go 语言中小数都是float64
	f2 := float32(1.23456)
	fmt.Printf("%T\n", f2) // 显式声明float32
	// f1 = f2 // float32 类型的值不能直接赋值给float64类型的变量
}
