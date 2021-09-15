package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 10
	fmt.Println("Value of a : ", a)
	fmt.Println("memory address of a : ", &a)
	var p *int = &a
	fmt.Println("Value of p : ", p)
	fmt.Println("deref of p : ", *p)
	fmt.Println("Type of p : ", reflect.TypeOf(p))
	fmt.Println("memory address of p : ", &p)
	var pointerToPointer **int = &p
	fmt.Println("Value of p2p : ", pointerToPointer)
	fmt.Println("deref of p2p : ", **pointerToPointer)
	fmt.Println("Type of p2p : ", reflect.TypeOf(pointerToPointer))
	fmt.Println("memory address of p2p : ", &pointerToPointer)
}
