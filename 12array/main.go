package main

import "fmt"

// 数组必须指定存放的元素的类型和长度
// 数组的长度是，数组类型的一部分
func main() {
	var a1 [3]bool // 可以理解为类型是 [3]bool
	var a2 [4]bool // 可以理解为类型是 [4]bool， 长度不同。
	fmt.Printf("bool: %T", true)
	fmt.Printf("a1: %T, a2: %T\n", a1, a2)

	// 数组的初始化
	// 如果不初始化： 默认元素都是零值（布尔值：false，浮点和整型都是0， 字符串：“”）
	fmt.Println(a1, a2)
	// 初始化方式1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	a1 = [3]bool{true, true}
	fmt.Println(a1)
	// 初始化方式2， 根据初始值自动推断数组长度是多少
	a10 := [...]int{0, 1, 2, 3, 4, 5, 5}
	fmt.Println(a10)
	// 初始化方式3, 根据索引来初始化
	a3 := [5]int{0: 1, 4: 2}
	fmt.Println(a3)

	// 数组的遍历
	cities := [...]string{"北京", "上海", "深圳"}
	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}
	for i, v := range cities {
		fmt.Println(i, v)
	}
	var a11 [2][2]int
	a11 = [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(a11)

	for i, v := range a11 {
		for j, k := range v {
			fmt.Println(i, j, k)
		}
	}
	// 数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[2] = 4
	b1[0] = 0
	fmt.Println(b1, b2)

	// 求和【1，3，5，7，8】
	// 找出数组中和为指定值的两个元素的下标，比如从数组【1，3，5，7，8】中找出和为8 的两个元素的下标分别为（0，3）和（1，2）
	a4 := [...]int{1, 3, 5, 7, 8}
	sum := 0
	cal := 8
	for _, v := range a4 {
		sum += v
	}
	fmt.Println(sum)
	for i := 0; i < len(a4)/2; i++ {
		for j := 0; j < len(a4); j++ {
			if (a4[i] + a4[j]) == cal {
				fmt.Printf("下标为%d和下标为%d的元素的和为%d\n", i, j, cal)
			}
		}
	}
	for i := 0; i < len(a4); i++ {
		for j := i + 1; j < len(a4); j++ {
			if (a4[i] + a4[j]) == cal {
				fmt.Printf("下标为%d和下标为%d的元素的和为%d\n", i, j, cal)
			}
		}
	}

}
