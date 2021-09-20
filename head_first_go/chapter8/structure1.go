package main

import "fmt"

type car struct {
	name     string
	topSpeed float64
}

func nitriBoost(c *car) {
	c.topSpeed += 50
}

func main() {
	var mustang car
	mustang.name = "Mustang Cobra"
	mustang.topSpeed = 225
	nitriBoost(&mustang)
	fmt.Println(mustang)
}
