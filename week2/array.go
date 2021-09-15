package main

import "fmt"

func main() {
	var myArray = [4]int{1, 2, 3, 4}
	fmt.Println(myArray)
	case1(myArray)
	fmt.Println("In main after case1 : ", myArray)
	case2(&myArray)
	fmt.Println("in main after case2 : ", myArray)
	var copyMyArray = myArray
	copyMyArray[0] = 50
	fmt.Println("my array is : ", myArray)
	fmt.Println("MyCopyArray is : ")
	for index, num := range copyMyArray {
		fmt.Println("index of ", index, " value ", num)
	}
}

func case1(arr [4]int) {
	arr[0] = 25
	fmt.Println("case1 value : ", arr)
}

func case2(arr *[4]int) {
	arr[0] = 15
	fmt.Println("case2 value : ", *arr)
}
