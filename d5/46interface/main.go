package main

import "fmt"

type cat struct{}
type dog struct{}
type person struct{}

func (c cat) speak() {
	fmt.Println("miaomiaomaio")
}
func (d dog) speak() {
	fmt.Println("wangwangwang")
}
func (p person) speak() {
	fmt.Println("aaaaa")
}

// 接口也是一种类型，它规定了，对象应该有的方法。

type speaker interface {
	speak()
}

//只要实现了speak 这个方法，都可以当作speaker
func da(x speaker) {
	x.speak()
}

func main() {
	var c1 cat
	var d1 dog
	var p1 person
	da(c1)
	da(d1)
	da(p1)

}
