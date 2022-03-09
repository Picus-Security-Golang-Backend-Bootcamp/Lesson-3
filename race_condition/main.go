package main

import (
	"fmt"
	"sync"
)

var counter = 0

func main() {

	wg := sync.WaitGroup{}

	wg.Add(2)
	mut := sync.Mutex{}
	rwmut := sync.RWMutex{}
	go func() {
		for i := 0; i < 10000; i++ {
			mut.Lock()
			counter++
			mut.Unlock()
		}

		wg.Done()
	}()

	go func() {
		for i := 0; i < 10000; i++ {
			mut.Lock()
			counter++
			mut.Unlock()

		}
		wg.Done()
	}()

	wg.Wait()

	fmt.Println(counter)
}
