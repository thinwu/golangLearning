package main

import "fmt"

func main() {
	n := 2
	switch n {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("default")
	}
	//变种 1
	switch j := 3; j {
	case 1, 3, 5, 7:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(j)
	}
	//变种 2
	j := 3
	switch {
	case j < 3:
		fmt.Println("aaa")
	case j > 3:
		fmt.Println("bbb")
	default:
		fmt.Println(j)
	}
	//变种3 fallthrough
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough //向下穿一个
	case s == "b":
		fmt.Println("b") //执行到这里，结束
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}

	//goto

	for i := 2; i < 5; i++ {
		if i == 3 {
			goto breakTag

		}
		fmt.Println(i)
	}
breakTag:
	fmt.Println("结束for 循环")

}
