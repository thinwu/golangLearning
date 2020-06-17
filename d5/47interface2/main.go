// the difference between value receiver and pointer receiver

package main

import "fmt"

type cat struct{}
type dog struct{}
type person struct {
	name string
}

func (c cat) speak() {
	fmt.Println("miaomiaomaio")
}
func (d dog) speak() {
	fmt.Println("wangwangwang")
}
func (p person) speak() {
	fmt.Println("aaaaa")
}

func (c *cat) move() {
	fmt.Println("cat run")
}
func (d *dog) move() {
	fmt.Println("dog run")
}
func (p *person) move() {
	fmt.Println("man run")
}

type mover interface {
	move()
}
type speaker interface {
	speak()
}

func main() {
	var p1 = person{
		name: "bbb",
	}
	var s1 speaker // define a speaker type, s1
	s1 = p1
	fmt.Println(s1)

	var m1 mover
	m1 = &p1 // the difference
	fmt.Println(m1)
}
