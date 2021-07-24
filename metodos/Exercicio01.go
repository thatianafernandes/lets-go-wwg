package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	raio float64
}

func (c Circulo) Area() float64 {
	return  math.Pi * math.Pow(c.raio , 2)
}

func (c Circulo) Perimetro() float64 {
	return  2 * math.Pi * c.raio
}

func main() {

	circulo := Circulo {
		raio : 6,
	}

	fmt.Println(circulo.Area())
	fmt.Println(circulo.Perimetro())
}