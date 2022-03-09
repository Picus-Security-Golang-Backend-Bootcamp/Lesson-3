package main

import (
	"fmt"
	"time"
)

func main() {

	//runtime.GOMAXPROCS(1)
	go func(i int) {
		for letter := 'c'; letter < 'c'+int32(i); letter++ {
			fmt.Printf("%c ", letter)
		}
	}(250)

	go func(i int) {
		for j := 0; j < i; j++ {
			fmt.Printf("%d ", j)
			//time.Sleep(time.Millisecond * 5)
		}
	}(300)

	time.Sleep(time.Second * 10)
}
