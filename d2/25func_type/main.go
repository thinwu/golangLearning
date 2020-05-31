package main

import "fmt"

// 函数类型

func f1() {
	fmt.Println("Hello")
}
func f2() int {
	return 2
}

// 函数也可以作为参数传

func f3(x func() int) int {
	return x()
}

// 函数还可以作为返回值

func f4() func(int, int) int {
	ret := func(a, b int) int {
		return a + b
	}
	return ret
}

func main() {
	a := f1
	fmt.Printf("%T\n", a)
	b := f2
	fmt.Printf("%T\n", b)
	fmt.Println(f3(b))
	a1 := f4()
	fmt.Println(a1(1, 2))

}
