package main

import "fmt"

// 函数定义，func [FuncName](signatures...)(return values){}

func sum(x int, y int) (z int) {
	return x + y
}

// 没有返回值
func f1(x int, y int) {
	fmt.Println(x + y)
}

// 没有参数，返回值

func f2() {
	fmt.Println("f2")
}

// 没有参数，有返回值, 返回值不命名
func f3() int {
	return 3
}

// 返回值可以命名也可以不命名
// 命名的返回值就相当于在函数中声明了一个变量
// 可以在函数体内使用返回值
func f4(x int, y int) (ret int) {
	ret = x + y
	return // 使用命名返回值可以直接使用return， 后面省略
}

// 多个返回值
func f5() (int, string) {
	return 1, "shake"
}

// 参数的类型简写, 多个连续参数，类型一样， 非最后一个的参数类型省略
func f6(x, y int, m, n string, i, j bool) int {
	return x + y
}

// 可变长参数， 必须放在函数参数的最后
func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y) // 是个切片
}

func main() {
	fmt.Println(sum(2, 4))
	s, v := f5()
	fmt.Println(s, v)
	f7("kong")
	f7("shanghe", 1, 2, 3, 4)
	f7("beijing", 23, 421)
}
