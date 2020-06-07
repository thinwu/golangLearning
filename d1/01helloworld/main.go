package main

import "fmt"

// 全局变量，声明后可不使用
// 函数内变量，声明后必须使用，返回值若不想用，可以用 _ 接收。_ 为
// var name string
// var age int
// var isOk bool

var (
	name string // ""
	age  int    // 0
	isOk bool   // false
)

//常量一般定义在全局
const pi = 3.1415926

const (
	statusOk = 200
	notFound = 404
)

// 批量声明时，如果一行声明后没有复制，默认使用上一行
const (
	n1 = 100
	n2
	n3
)

//iota: 枚举
const (
	a1 = iota //重置为0
	a2
	a3
)

const (
	b1 = iota // 0
	b2        // 1
	_         //2， _ 哑元变量，这个变量不会分配内存，直接舍弃，使用不到，用于接收不想用的返回值等情况。
	b3        //3
)

//插队
const (
	c1 = iota //0
	c2 = 100
	c3        // 100
	c4 = iota // 3
)

//多个常量声明在一行
const (
	d1, d2 = iota + 1, iota + 2 //d1:1, d2:2 iota = 0
	d3, d4 = iota + 1, iota + 2 // d3:2,d4 :3, iota = 1
)

// 定义数量级

const (
	_  = iota             // 0 舍去
	kb = 1 << (10 * iota) // iota = 1， 10000000000
	mb = 1 << (10 * iota) // iota = 2， 100000000000000000000
	gb = 1 << (10 * iota)
	tb = 1 << (10 * iota)
	pb = 1 << (10 * iota)
)

func main() {
	name = "理想"
	age = 16
	isOk = true
	fmt.Println("helloworld")
	fmt.Println("name:" + name)
	var s1 string = "whb"
	fmt.Println(s1)
	//类型推导
	var s2 = "20"
	fmt.Printf(s2)
	// 简短变量声明, 只能在函数里用
	s3 := "haha"
	fmt.Println(s3)
	fmt.Println(c4)
	fmt.Println(statusOk)

	// go build 直接编译操作系统平台的可执行文件， main.go 为程序主入口

	//交叉编译， 平台A编译平台B的程序
	// win =》 linux or Mac
	/*
		CMD中
		SET CGO_ENABLED=0 	// 禁用 CGO
		SET GOOS=linux 		// 目标平台linux
		SET GOARCH=amd64 	//目标处理器amd64
		go build  // 编译

		SET CGO_ENABLED=0  // 禁用 CGO
		SET GOOS=darwin    // 目标平台Mac
		SET GOARCH=amd64   // 目标处理器amd64
		go build  // 编译
	*/

	// linux or Mac 下交叉编译， 直接编译的话，go build 即可
	/*
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
		CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build
		CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
	*/
	// go run main.go => 编译在临时文件夹，再执行。
}
