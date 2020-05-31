package main

import "fmt"

// append () 为切片追加元素
func main() {
	s1 := []string{"北京", "深圳", "上海"}
	fmt.Printf("s1=%v, len(s1):%d,cap(s1):%d\n", s1, len(s1), cap(s1))
	// 调用append函数，用原来的切片变量接收返回值
	s1 = append(s1, "广州") // 重新换底层数组
	fmt.Printf("s1=%v, len(s1):%d,cap(s1):%d\n", s1, len(s1), cap(s1))
	s2 := append(s1, "重庆") // 重新换底层数组
	fmt.Printf("s2=%v, len(s2):%d,cap(s2):%d\n", s2, len(s2), cap(s2))

	s2 = append(s1, "杭州", "成都")
	fmt.Printf("s2=%v, len(s2):%d,cap(s2):%d\n", s2, len(s2), cap(s2))
	s3 := []string{"wuhan", "xi'an", "suzhou"}
	s2 = append(s2, s3...) //... 表示拆开切片，回到值类型
	fmt.Printf("s2=%v, len(s2):%d,cap(s2):%d\n", s2, len(s2), cap(s2))

	var s4 []int       // 没有初始化
	s4 = append(s4, 1) // append 可以对没有初始化的切片做操作
	fmt.Println(s4)

}
