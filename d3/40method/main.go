package main

import "fmt"

// 方法，是作用域特定类型的函数
// 标识符： 变量名，函数名，类型名，方法名
// Go语言中如果标识符首字母是大写的，就是共有的（public）
// 大写的标识符都需要写注释注释格式 标识符 空格 解释=

//比如

// Dog 这个是一个示例
type Dog struct {
	name string
}

func newDog(name string) *Dog {
	return &Dog{
		name: name,
	}
}

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

// 方法
// 接收者表示的是调用该方法的具体类型变量，多用类型首字母小写
// 接收者写在函数名前
// func (接收者变量名，接收者类型) 方法名(参数列表) (返回值){函数体}
// 下面这个例子是值接收者
func (d Dog) wang() {
	fmt.Printf("%s: wangwangwang\n", d.name)
}

func (p person) guonian() {
	p.age++
}

//使用指针接收者
// 1. 当修改接收者中的值
// 2. 接收者是copy代价比较大的对象
// 3. 保证一致性，如果某个方法是用指针类型，那么以后就用指针接受
func (p *person) pGuonian() {
	p.age++
}

func main() {
	d := newDog("aaa")
	d.wang()
	p := newPerson("bb", 18)
	fmt.Printf("%v\n", p.age)
	p.guonian() // 值类型的接收者。传的copy，改不了结构体里的值
	fmt.Printf("%v\n", p.age)
	(*p).pGuonian()
	fmt.Printf("%v\n", p.age)

}
