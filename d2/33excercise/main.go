/*50枚金币，分给：Mattew, Sarah, Augustus, Heidi, Emilie, Peter, Giana, Adriano, Aaron, Elizabeth
	分配规则如下：
	名字中每包含1个e 或者E 分1枚
	名字中每包含1个i 或者I 分2枚
	名字中每包含1个o 或者O 分3枚
	名字中每包含1个u 或者U 分4枚
计算每个用户分到多少金币，以及最后还有多少
*/
package main

import "fmt"

var (
	coins = 50
	users = []string{
		"Mattew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() int {
	coinUsed := 0
	for _, name := range users {
		coin := 0
		for _, c := range name {
			switch c {
			case 'e', 'E':
				coin++
			case 'i', 'I':
				coin += 2
			case 'o', 'O':
				coin += 3
			case 'u', 'U':
				coin += 4
			}
		}
		distribution[name] = coin
		coinUsed += coin
	}
	fmt.Printf("%v\n", distribution)
	return coins - coinUsed
}
func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
}
