package main

import "fmt"

func main() {
	mySlice := make([]int, 5)
	for i := 0; i < len(mySlice); i++ {
		mySlice[i] = (i + 1) * 10
	}
	fmt.Println("mySlice is: ", mySlice)

	newSlice := mySlice[:]
	fmt.Println(newSlice)

	newSlice1 := mySlice[2:]
	fmt.Println(newSlice1)

	newSlice2 := mySlice[:2]
	fmt.Println(newSlice2)

	newSlice3 := mySlice[2:4]
	fmt.Println(newSlice3)

	mySlice = append(mySlice[:2], mySlice[3:]...)
	fmt.Println("mySlice is: ", mySlice)

}
