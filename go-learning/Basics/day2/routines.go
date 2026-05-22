package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go printNumbers()

	time.Sleep(time.Second * 2)
}