package main

import (
	"fmt"
	"head_first_go/chapter10/geo1"
)

func main() {
	coordinates := geo1.Coordinates{}
	coordinates.SetLatitude(37.42)
	coordinates.SetLongitude(-122.08)
	fmt.Println(coordinates)
}
