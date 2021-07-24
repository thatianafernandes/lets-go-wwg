package main

import "fmt"

type Apartamento struct {
	Proprietario string
	Numero int
	PossuiVagaGaragem bool
}

func criaPredio() []Apartamento {
	ap1 := Apartamento{
		Numero:              1301,
		Proprietario:  "Thalía",
		PossuiVagaGaragem: true,
	}

	ap2 := Apartamento{
		Numero:              1302,
		Proprietario:  "Shakira",
		PossuiVagaGaragem: false,
	}

	ap3 := Apartamento{
		Numero:              1401,
		Proprietario:  "Anitta",
		PossuiVagaGaragem: true,
	}

	apartamentos := []Apartamento{ap1, ap2, ap3}
	return apartamentos;
}

func main() {
	predio :=  criaPredio()

	fmt.Println(predio)
}