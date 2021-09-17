package main

import (
	"fmt"
)

type myError struct {
	s string
}

func (err *myError) Error() string {
	return err.s
}

func main() {
	quotient, err := divide(4, 2)
	if err != nil {
		fmt.Println("Error occured : ", err)
		return
	}
	fmt.Println("Result is : ", quotient)
}

func divide(numerator, denominator int) (int, error) {
	if denominator == 0 {
		return 0, &myError{s: "cannot divide by zero"}
	}
	return numerator / denominator, nil
}
