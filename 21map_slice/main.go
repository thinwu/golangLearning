package main

import (
	"fmt"
	"strings"
)

// map and slice

func main() {

	// 切片和map 都要做初始化

	//元素类型为map的切片, 预估一下容量
	//s1 := make([]map[int]string, 0, 10)
	//s1[0][100] = "A" // => index of range, make 的切片里面长度是0。s1[0] == nil,
	//就算长度不是0， 比如make([]map[int]string, 1, 10) 那么s1[0][100] 这里的map也没有初始化， s1[0][100] = nil

	//fmt.Println(s1)

	// 救一下
	s1 := make([]map[int]string, 0, 10)
	m1 := make(map[int]string, 1) // 创建一个新map
	m1[0] = "jiuyixia"            // 赋值
	s1 = append(s1, m1)           // 追加进s1 的空位
	fmt.Println(s1)

	// 对切片长度已经有定义，切片划分了内存空间了。但是map还没有初始化
	s2 := make([]map[int]string, 10, 10)
	s2[0] = make(map[int]string, 1)
	s2[0][10] = "shake"
	fmt.Println(s2)

	//元素类型为切片的map

	var m2 = make(map[string][]int, 5)
	m2["北京"] = []int{1, 2, 3, 4, 5, 6, 4}
	fmt.Println(m2)

	w1 := "how do you do"
	m3 := make(map[string]int, 10)

	for _, v := range strings.Split(w1, " ") {
		m3[string(v)]++
	}
	fmt.Println(m3)
}
