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
	go writeToChannel(ch)
	go readToChannel(ch)
	wg.Wait()
}

func writeToChannel(ch chan int) {
	time.Sleep(time.Second)
	fmt.Println("Starting to Write")
	ch <- 10
	ch <- 8
	fmt.Println("Data has been written to channel")
	wg.Done()
}

func readToChannel(ch chan int) {
	fmt.Println("Starting to Read")
	value := <-ch
	value1 := <-ch
	fmt.Println("Data of our channel is: ", value, value1)
	wg.Done()
}
