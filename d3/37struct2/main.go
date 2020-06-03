package main

import "fmt"

type person struct {
	gender, name string
}

func f(x person) {
	x.gender = "nv"
}
func f2(x *person) {
	// 1.(*x).gender = "nv" // 根据内存地址找到那个原来的变量，修改的就是原来的变量
	x.gender = "nv" // 语法糖。如果x是指针，因为go 是不能对指针操作。所以肯定找的是指针对应的变量
}
func main() {
	var p person
	p.gender = "nan"
	p.name = "zhouling"
	f(p)
	fmt.Println(p)
	f2(&p)
	fmt.Println(p)
}
