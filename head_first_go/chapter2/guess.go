package main

import (
	"fmt"
	"math/rand"
)

func main() {
	target := rand.Intn(100) + 1
	fmt.Println("Enter your target number from 1 to 100")
	var number, i int
	for i = 0; i <= 10; i++ {
		_, err := fmt.Scanln(&number)
		checkError(err)
		if target == number {
			fmt.Println("Good Job! you guessed it!")
			break
		} else if number < target {
			fmt.Println("Oops your guess was LOW")
		} else {
			fmt.Println("Oops your guess was HIGH")

		}
	}
	if i == 11 {
		fmt.Println("Soory, you didn't guess my number")
		fmt.Println("it was taget : ", target)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
