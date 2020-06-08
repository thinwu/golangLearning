package main

import "fmt"

// nested struct

type person struct {
	age  int
	base base
}
type company struct {
	base base
}
type company2 struct {
	test string
	base // nested anonymous struct, only one nested anonymous struct to avoid possible field conflicts
}

type base struct {
	name     string
	province string
	city     string
}

func main() {
	p1 := person{
		age: 37,
		base: base{
			name:     "Daniel",
			province: "shanghai",
			city:     "city",
		},
	}
	fmt.Println(p1.base.name)
	// fmt.Println(p1.name) 无法访问

	c1 := company2{
		test: "aaa",
		base: base{
			name:     "company2",
			province: "shanghai",
			city:     "shanghai",
		},
	}
	fmt.Println(c1.name) // accessible, will first look self struct, then jump into the nested anonymous struct
}
