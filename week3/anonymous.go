package main

import "fmt"

func main() {
	func() {
		fmt.Println("inside anonymous function")
	}()
	anonyFunc()
}

func anonyFunc() {
	var number = func(s string) int {
		fmt.Println(s, "inside non main anonymous function")
		return 10
	}("Hello Anupam")
	fmt.Println(number)
}
