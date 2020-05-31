package main

import "fmt"

//匿名函数， 一般都用在函数内部

// var f1 = func(x, y int) {
// 	fmt.Println(x + y)
// }

func main() {
	//f1(10, 20)
	//函数内部没有办法声明带名字的函数
	f1 := func(x, y int) {
		fmt.Println(x + y)
	}
	f1(10, 20)
	// 如果只是调用一次的函数，还可以简写成立即执行函数， 最后加括号(),表示调用
	func() {
		fmt.Println("helloworld")
	}()

	func(x, y int) { // 函数参数定义
		fmt.Println(x + y)
	}(100, 200) //函数定义
	//函数调用, 函数参数传递

	/*func(x, y int) { // 函数参数定义
		fmt.Println("helloworld")
	} //函数定义
	(100,200) //函数调用, 函数参数传递
	*/
}
