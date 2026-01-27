package main

import (
	"fmt"
	"sync"
)

func Go() {
	var wg sync.WaitGroup
	wg.Add(2)
	var wg2 sync.WaitGroup
	wg2.Add(1)
	channel := make(chan int)
	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j <= 10; j++ {
				channel <- j
			}
		}()
	}
	go func() {
		sum := 0
		defer wg2.Done()
		for {
			out, ok := <-channel
			if !ok {
				return
			}
			sum += out
			fmt.Println(sum)
		}
	}()

	wg.Wait()
	close(channel)
	wg2.Wait()

}

func main() {
	fmt.Println("Chapter 8")
	Go()

}
