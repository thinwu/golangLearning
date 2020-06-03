package main

import "fmt"

//构造函数

type person struct {
	name string
	age  int
}

// 构造函数，约定用new开头
//返回的是结构体，还是结构体指针，是需要思考的。
// 返回的是值，内存中就有一次copy操作。因为go 是传值的。
// 当结构体比较大的时候，尽量使用结构体指针，减少程序的运行内存开销。
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func main() {
	p1 := newPerson("Daniel", 18)
	p2 := newPerson("Cyrus", 10)
	fmt.Println(p1, p2)
}
