package main

import (
	"fmt"
	"time"
)

func getCustomerId(ch chan int) {
	time.Sleep(5 * time.Second)
	ch <- 20
}

func main() {
	ch := make(chan int)

	go getCustomerId(ch)

	for {
		time.Sleep(1 * time.Second)
		select {
		case a := <-ch:
			fmt.Println(a)
			return
		default:
			fmt.Println("No data")
		}

	}
}
