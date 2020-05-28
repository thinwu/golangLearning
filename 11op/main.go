package main

import "fmt"

func main() {
	var (
		a = 5
		b = 2
	)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++ // go里面作为单独的语句， 不能放在等号的右边 ==》 a=a+1
	b-- // ==》 b = b-1

	//位运算符：针对的是二进制数， 可用于权限管理
	// 5 的二进制表示 101
	// 2 的而精致表示 10 ==》 010

	// & 按位 与（两位均为1 才为1）
	fmt.Println(5 & 2) // 000 =>0
	// | 按位 或（两位有一个为1才为1）
	fmt.Println(5 | 2) // 111 =>7(10 进制打印结果)
	// ^ 按位异或（两位不一样才为1）
	fmt.Println(5 ^ 2) // 111
	// << 将二进制位左移指定位数
	fmt.Println(5 << 1)  // 101 =》 1010 == 十进制的10
	fmt.Println(1 << 10) // 10000000000 == 十进制的1024 可以用来计算容量
	// >> 将二进制位右移动指定位数
	fmt.Println(5 >> 1) // 101 => 10 == 十进制2
	fmt.Println(5 >> 3) // 101 => 0 == 十进制0

	//var m = int8(1)
	//fmt.Println(m << 10) // 左移了10位但是只能放8位，最后也只能存8位

	var x int
	x = 10
	x += 1 // => x++
	x -= 1 // => x--
	x *= 2
	x /= 2
	x %= 2
	x &= 2
	x |= 2
	x ^= 2
	x >>= 2
	x <<= 2
	fmt.Println(x)
}
