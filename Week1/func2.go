package main

import "fmt"

func main() {
	userNumber := addOne(19)
	fmt.Println(userNumber)
}

func addOne(num int) int {
	return num + 1
}
