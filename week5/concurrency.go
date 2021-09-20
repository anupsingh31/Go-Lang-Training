package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	runtime.GOMAXPROCS(-1)
	wg.Add(2)
	go delayIteration1()
	go delayIteration2()
	wg.Wait()
	//fmt.Scanln()
}

func delayIteration1() {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("In 1 second : ", i)
	}
	wg.Done()
}

func delayIteration2() {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("In 2 second : ", i)
	}
	wg.Done()
}
