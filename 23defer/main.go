package main

import "fmt"

func deferDemo() {
	fmt.Println("Start")
	defer fmt.Println("Defer1") // 将它后面的语句延迟到，函数即将返回时执行
	fmt.Println("end")          // 这里结束后即将返回，defer开始执行
	// 可以用于函数结束前释放资源，如：关闭连接，文件，数据库访问，等，写在开始，以免忘记，但又在最后执行。

}

// Go语言中函数的return 不是原子操作，在底层分为两步来执行
// 第一步：返回值赋值
// 第二步：真正的RET操作
// 函数中如果存在defer， 那么defer 执行的是在第一步和第二步之间

func f1() int {
	x := 5
	defer func() { x++ }()
	return x // 返回值等于x=5， x++， 返回 5，因为没有返回x++
}

func f2() (x int) {
	defer func() { x++ }()
	return 5 // 6， 1. 返回值x=5, x++=6, 返回x，=>6
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x // 5. => y = 5
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func main() {
	deferDemo()
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
