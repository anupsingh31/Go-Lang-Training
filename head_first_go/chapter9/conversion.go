package main

import "fmt"

type Litres float64
type Millilitres float64

func (l Litres) toMillilitres() Millilitres {
	return Millilitres(l * 1000)
}

func (m Millilitres) toLitres() Litres {
	return Litres(m / 1000)
}

func main() {
	l := Litres(3)
	fmt.Println(l, " Litres is ", l.toMillilitres(), " millilitres")
	ml := Millilitres(500)
	fmt.Println(ml, " millilitres", ml.toLitres(), " litres ")

}
