// type assert， 类型断言

package main

import "fmt"

func assert(a interface{}) {
	str, ok := a.(string)
	if !ok {
		fmt.Println("not String")
	} else {
		fmt.Println(str)
	}

}
func assert2(a interface{}) {
	switch a.(type) {
	case string:
		fmt.Printf("string %v", a.(string))
		break
	case int:

	}

}
func main() {
	assert(100) // runtime error, if in assert there is no if-ok
}
