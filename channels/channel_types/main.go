package main

import (
	"fmt"
	"time"
)

func writeMessage(ch chan string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Wrote %d to channel\n", i)
		ch <- fmt.Sprintf("Hello, %d", i)

	}

	close(ch)
}

func main() {
	//ch := make(chan string) //unbuffered channel

	ch := make(chan string, 5) //buffered channel

	go writeMessage(ch)
	go readMessage(ch)

}

func readMessage(ch chan string) {
	for value := range ch {
		fmt.Println(value)
		time.Sleep(time.Second)
	}
}
