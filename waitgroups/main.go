package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(value int) {
			fmt.Println(value)
			wg.Done()
		}(i)

		wg.Wait()
	}

}
