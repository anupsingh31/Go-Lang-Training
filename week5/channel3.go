package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(2)
	ch := make(chan int, 2)
	ch2 := make(chan struct{})
	go writeToChannel(ch)
	go readToChannel(ch)
	wg.Wait()
	close(ch)
	go func() {
		ch2 <- struct{}{}
	}()

	for i := range ch {
		fmt.Println("Inside range ", i)
	}
	for i := 0; i < 2; i++ {
		select {
		case chan1, ok := <-ch:
			fmt.Println("channel 1 has data ", chan1, ok)
		case chan2, ok := <-ch2:
			fmt.Println("channel 2 has data ", chan2, ok)

		}
	}
}

func writeToChannel(ch chan<- int) {
	time.Sleep(time.Second)
	fmt.Println("Starting to Write")
	ch <- 10
	ch <- 8
	ch <- 6
	fmt.Println("Data has been written to channel")
	wg.Done()
}

func readToChannel(ch <-chan int) {
	fmt.Println("Starting to Read")
	value := <-ch
	value1 := <-ch
	fmt.Println("Data of our channel is: ", value, value1)
	wg.Done()
}
