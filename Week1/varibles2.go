package main

import (
	"fmt"
	"runtime"
)

var name string

func main() {
	var rollNo int
	rollNo = 1
	var (
		firstName, lastName, flag = "Anupam", "Singh", true
	)
	test := 20

	fmt.Println(rollNo)
	fmt.Println(firstName, lastName, flag)
	fmt.Println(test)
	fmt.Printf("%T", test)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
