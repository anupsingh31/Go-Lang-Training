package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myArray = [4]string{"hello", "from", "where", "are"}
	fmt.Println(reflect.TypeOf(myArray))
	case1(myArray)
	for index, val := range myArray {
		fmt.Println("Index is : ", index, "Value is : ", val)
	}
	case2(&myArray)
	for index, val := range myArray {
		fmt.Println("Index is : ", index, "Value is : ", val)
	}

}

func case1(arr [4]string) {
	arr[0] = "hey"
	fmt.Println("In case1 : ", arr)

}

func case2(arr *[4]string) {
	arr[0] = "hey"
	fmt.Println("In case2 (deref) : ", *arr)
	fmt.Println("In case2 : ", arr)
}
