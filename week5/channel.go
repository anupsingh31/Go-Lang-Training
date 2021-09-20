package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go writeToChannel(ch)
	value := <-ch
	fmt.Println("Data of our channel is: ", value)

}

func writeToChannel(ch chan int) {
	time.Sleep(time.Second)
	fmt.Println("Inside Go routine")
	ch <- 10
	fmt.Println("Data has been written to channel")
}
