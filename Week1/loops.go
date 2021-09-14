package main

import "fmt"

func main() {
	i := 0
	for {
		fmt.Println("Hello")
		i += 5
		if i == 20 {
			break
		}

	}
	for j := 0; j < 10; j++ {
		fmt.Println("Hello", j)
	}

	sum := 0
	for ; sum < 5; sum++ {
		fmt.Println(sum)
	}
}
