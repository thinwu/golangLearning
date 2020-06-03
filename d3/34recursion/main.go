package main

import "fmt"

// 递归： 自己调用自己
//适合处理问题相同，规模越来越小的场景

func f(n uint64) uint64 {
	if n <= 0 {
		return 1 // 确定的结束条件，否则是死循环
	}
	return n * f(n-1)
}

// 上台阶
// n 个台阶，一次可以走一步，也可以走两步，有多少种走法。

func plat(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return plat(n-1) + plat(n-2)
}

func main() {
	fmt.Println(f(7))
}
