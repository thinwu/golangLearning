package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	timestamp1 := now.Unix()
	fmt.Println(timestamp1)
	timestamp2 := now.UnixNano()
	fmt.Println(timestamp2)
	tmp := time.Unix(timestamp1, 0)
	fmt.Println(tmp)
	fmt.Println(tmp.Add(24 * time.Hour))
	fmt.Println(tmp.Sub(now))
	// ticker
	// timer := time.Tick(time.Second)
	// for t := range timer {
	// 	fmt.Println(t)
	// }
	// using the time of GoLang created 2006 1 2 15 04 05 -> Y M D H m S

	// time format, from time to string yyyy-mm-dd HH:MM:SS pm
	fmt.Println(now.Format("2006-01-02 03:04:05 PM"))
	// placeholder yyy-mm-dd HH:MM:SS.[millionSeconds placeholder] pm
	fmt.Println(now.Format("006-01-02: 15:04:05.0000000 "))

	// from string to time
	timeObj, err := time.Parse("2006-01-02", "2000-08-03")
	if err != nil {
		fmt.Printf("convert time error, %v\n", err)
		return
	}
	fmt.Println(timeObj)

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(loc)
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-07-02 15:06:05", loc)
	fmt.Println(t)
}
