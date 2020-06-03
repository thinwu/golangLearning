package main

import "fmt"

func main() {
	// & 取地址
	// * 根据地址取值
	n := 18
	fmt.Println(&n)
	p := &n
	fmt.Printf("%T\n", p)
	fmt.Printf("&n=%p\n", &n) // n的内存地址
	fmt.Printf("p=%p\n", p)   // p这个指针的值，就是n的地址
	fmt.Printf("&p=%p\n", &p) // p的地址
	m := *p
	fmt.Println(m)
	fmt.Printf("%T\n", m)

	// var a *int // nil
	// *a = 100   // nil = 100 => runtime exception
	// fmt.Println(*a)

	// new 函数申请一个内存地址
	var a = new(int) // 获得一个内存地址
	fmt.Println(a)
	fmt.Println(*a) // 对应类型的默认值
	*a = 100
	fmt.Println(*a)
}
