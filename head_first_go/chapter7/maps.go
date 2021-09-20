package main

import "fmt"

func main() {
	jewelry := make(map[string]float64)
	jewelry["necklace"] = 89.99
	jewelry["earring"] = 79.99
	clothing := map[string]float64{"paint": 59.44, "shirt": 67.77}
	fmt.Println("Necklace : ", jewelry["necklace"])
	fmt.Println("Earring : ", jewelry["earring"])
	fmt.Println("Shirt : ", clothing["shirt"])
	fmt.Println("Paint : ", clothing["paint"])

}
