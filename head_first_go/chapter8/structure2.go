package main

import "fmt"

type part struct {
	discription string
	count       int
}

func doublePack(p *part) {
	p.count *= 2
}

func main() {
	var fuses part
	fuses.discription = "Fuses"
	fuses.count = 5
	doublePack(&fuses)
	fmt.Println(fuses)
}
