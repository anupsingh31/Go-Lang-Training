package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 6 {
			continue
		}
		if i == 8 {
			break
		}
		fmt.Println("Iteratot ", i)
	}
	fmt.Println("We are outside the loop")
}
