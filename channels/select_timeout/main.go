package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	check(ctx)

	fmt.Println("Sdfsdfsdfsdf")

}

func check(ctx context.Context) {
	ticker := time.NewTicker(time.Second * 3)
	for {
		select {
		case <-ticker.C:
			writeMessage()
		case <-ctx.Done():
			fmt.Println("No data")
			return
		}
	}
}

func writeMessage() {
	log.Println("Hello world")
}
