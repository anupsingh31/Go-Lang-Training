package main

import (
	"fmt"
	"head_first_go/chapter8/geo"
)

func main() {
	location := geo.Landmark{}
	location.Name = "Thane"
	location.Latitude = 123.45
	location.Longitude = 12.44
	fmt.Println(location)
}
