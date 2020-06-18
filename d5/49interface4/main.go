// empty interface

package main

import "fmt"

// interface{}
// all types have implemented this interface
// if you have a function takes what ever types as a param, you should use empty interface

func main() {
	m1 := make(map[string]interface{}, 10) // a map with int type as key, and any type as value
	fmt.Println(m1)
	m1["name"] = "aaa"
	m1["age"] = 100
	m1["merried"] = true
	// three types value
	m1["hobby"] = [...]string{"sing", "dance", "rap"}
	fmt.Println(m1)
	show(true, false)
	show(m1)
}

func show(a ...interface{}) {
	fmt.Println(a...)
}
