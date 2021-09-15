package main

import "fmt"

func main() {
	mySlice := make([]int, 5)
	for i := 0; i < len(mySlice); i++ {
		mySlice[i] = (i + 1) * 10
	}
	fmt.Println(mySlice)

	newSlice := mySlice
	newSlice[0] = 90
	fmt.Println("mySlice is: ", mySlice)
	fmt.Println("newSlice is: ", newSlice)

}
