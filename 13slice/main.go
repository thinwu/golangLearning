package main

import "fmt"

func main() {
	// 定义切片
	var s1 []int    // defint an int type slice
	var s2 []string // define an string type slice
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	s1 = []int{1, 2, 3}
	fmt.Println(s1 == nil)
	s2 = []string{"shanghe", "zhangjiang", "pingshancun"}
	fmt.Println(s1, s2)

	// length and capacity
	fmt.Printf("len(s1): %d cap(s1): %d\n", len(s1), cap(s1))
	fmt.Printf("len(s2): %d cap(s2): %d\n", len(s2), cap(s2))

	// getting slice from array
	a1 := [...]int{1, 3, 4, 5, 6, 7, 8}
	// includes left, exlucde right
	s3 := a1[0:4] // slicing based on an array start from 0, exclude the forth
	fmt.Println(s3)
	s3 = a1[1:5] //3, 4, 5, 6,
	fmt.Println(s3)
	s4 := a1[:4] // =>[0:4]
	s5 := a1[3:] // => [3: len(a1)]
	s6 := a1[:]  // => [0:len(a1)]
	fmt.Println(s4, s5, s6)
	// the capacity of a slice is the length of the start element to the end of its basic array.
	fmt.Printf("len(s4): %d cap(s4): %d\n", len(s4), cap(s4))
	fmt.Printf("len(s5): %d cap(s5): %d\n", len(s5), cap(s5))

	s7 := s5[1:]
	fmt.Printf("len(s7): %d cap(s7): %d\n", len(s7), cap(s7))
	s8 := s3[2:]
	fmt.Printf("len(s3): %d cap(s3): %d\n", len(s3), cap(s3))
	fmt.Printf("len(s8): %d cap(s8): %d\n", len(s8), cap(s8))
	// slice is a reference type, pointing to the basic array
	a1[4] = 1300 // changing the basic array value, the slice value will be changed as well
	fmt.Println("s8", s8)
	fmt.Println("s7", s7)
	fmt.Println("s6", s6)

}
