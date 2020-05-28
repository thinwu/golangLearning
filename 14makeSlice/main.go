package main

import "fmt"

func main() {
	s1 := make([]int, 5, 10)
	s2 := make([]int, 5)
	fmt.Printf("s1=%v, len(s1):%d, cap(s1):%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2):%d, cap(s2):%d\n", s2, len(s2), cap(s2))
	s3 := make([]int, 0, 10)
	fmt.Printf("s3=%v, len(s3):%d, cap(s3):%d\n", s3, len(s3), cap(s3))

	var s4 []int         // len(s1)=0; cap(s1)=0; s1 == nil
	s5 := []int{}        // len(s2)=0; cap(s2)=0; s2 != nil
	s6 := make([]int, 0) // len(s)=0; cap(s1)=0; s3 != nil
	fmt.Println(s4, s5, s6)
	fmt.Println(s4 == nil, s5 == nil, s6 == nil)
	// 判断切片是否为空，用len(s) == 0

	//切片的复制copy，共用低层数据
	s7 := []int{1, 2, 3}
	s8 := s7
	s7[0] = 100
	fmt.Println(s7)
	fmt.Println(s8)

	// 切片的遍历
	// 1.索引遍历
	for i := 0; i < len(s7); i++ {
		fmt.Println(s7[i])
	}
	for i, v := range s7 {
		fmt.Println(i, v)
	}
}
