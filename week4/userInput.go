package main

import (
	"fmt"
	"log"
)

func main() {
	// fmt.Println("Type something :")
	// reader := bufio.NewReader(os.Stdin)
	// input, _ := reader.ReadString('\n')
	// fmt.Println("Input is : ", input)

	fmt.Println("Type something :")
	var number int
	_, err := fmt.Scan(&number)
	if err != nil {
		log.Fatal("Error : ", err)
		// fmt.Println("error : ", err)
	}

	fmt.Println("number is : ", number)

}
