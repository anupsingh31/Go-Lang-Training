package main

import "fmt"

func main() {
	var myMap = make(map[string]int)
	myMap["a"] = 10
	myMap["b"] = 20
	myMap["c"] = 30
	delete(myMap, "ac")
	var num, ok = myMap["c"]
	fmt.Println("Num is : ", num)
	fmt.Println("ok values : ", ok)
	for key, value := range myMap {
		fmt.Println("Key is : ", key, "Values is : ", value)
	}

}
