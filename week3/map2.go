package main

import "fmt"

func main() {
	myMap := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}
	fmt.Println("Mymap is : ", myMap)
	newMap := checkMap(myMap)
	fmt.Println("newmap is : ", newMap)

}

func checkMap(myMap map[int]string) map[int]string {
	myMap[1] = "changed"
	return myMap
}
