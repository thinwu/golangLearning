package main

import "fmt"

func main() {
	a1 := []int{1, 3, 5} // 切片
	a2 := a1
	var a3 []int // nil, 没有内存不能copy
	copy(a3, a1)
	fmt.Println(a1, a2, a3)

	a3 = make([]int, 3, 3)
	copy(a3, a1) // 创建了一个新的低层数组
	fmt.Println(a1, a2, a3)
	a1[0] = 100
	fmt.Println(a1, a2, a3)

	// del

	a1 = append(a1[:1], a1[2:]...) // append([100],[5])
	fmt.Println(a1)
	fmt.Println(cap(a1))

	x1 := [...]int{1, 3, 5}    // 数组
	s1 := x1[:]                // 框住了x1头尾的切片，切片本身不存值
	fmt.Printf("%p\n", &s1[0]) // 打印地址
	fmt.Println(s1, len(s1), cap(s1))
	s1 = append(s1[:1], s1[2:]...) // 修改了低层数组

	fmt.Println(x1)
	fmt.Printf("%p\n", &s1[0]) // 打印地址

}
