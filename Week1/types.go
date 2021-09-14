package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf("hello"))
	fmt.Println(reflect.TypeOf(1200))
	fmt.Println(reflect.TypeOf(12.00))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("49"))
}
