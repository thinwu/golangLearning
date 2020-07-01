package main

import (
	"log"
	"os"
)

func main() {
	logFi, _ := os.OpenFile("./xx.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
	v := "a general log"
	log.SetOutput(logFi)
	log.Printf("something with printf, %v\n", v)
	log.Panicf("something with Panicf, %v\n", v)
	log.Fatalf("something with Fatalf, %v\n", v)
}
