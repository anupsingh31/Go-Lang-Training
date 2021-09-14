package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().Hour()
	fmt.Println(now)
	if now <= 12 {
		fmt.Println("Good Morning")
	} else if now < 17 {
		fmt.Println("Good After Noon")
	} else if now < 22 {
		fmt.Println("Good Evening")
	} else {
		fmt.Println("Good night")
	}
}
