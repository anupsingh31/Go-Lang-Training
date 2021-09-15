package main

import "fmt"

func main() {
	var mySlice1 []int
	fmt.Println(mySlice1)
	fmt.Println(mySlice1 == nil)
	mySlice := make([]int, 5)
	for i := 0; i < len(mySlice); i++ {
		mySlice[i] = (i + 1) * 10
	}
	fmt.Println("mySlice is: ", mySlice)
	fmt.Println("mySlice length is: ", len(mySlice))
	fmt.Println("mySlice capacity is: ", cap(mySlice))
	fmt.Println("mySlice first address is: ", &mySlice[0])

	mySlice = append(mySlice, 60, 70)
	fmt.Println("mySlice is: ", mySlice)
	fmt.Println("mySlice length is: ", len(mySlice))
	fmt.Println("mySlice capacity is: ", cap(mySlice))
	fmt.Println("mySlice first address is: ", &mySlice[0])

	newSlice := append(mySlice, 80, 90, 100, 110)
	fmt.Println("newSlice is: ", newSlice)
	fmt.Println("newSlice length is: ", len(newSlice))
	fmt.Println("newSlice capacity is: ", cap(newSlice))
	fmt.Println("newSlice first address is: ", &newSlice[0])

}
