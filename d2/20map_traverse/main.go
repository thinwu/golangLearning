package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	//time.Now() 拿到当前时间，UnixNano()拿到当前时间的纳秒数
	rand.Seed(time.Now().UnixNano()) // 初始化随机数种子
	var scoreMap = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) // 右对齐，总长2位
		value := rand.Intn(100)
		scoreMap[key] = value
	}
	// 取出map中所有的key 存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	// 对切片进行排序
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}
