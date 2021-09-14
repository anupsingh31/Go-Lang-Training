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
	for j := 10; j > 0; j-- {
		fmt.Println("Hello", j)
	}

	sum := 10
	for sum > 0 {
		fmt.Println(sum)
		sum -= 2
	}
}
