package main

import (
	"fmt"
	"time"
)

type Venda struct {
	Valor int
	Nome string
	Data time.Time
}

func totalVendido(vendasMiojo []Venda) map[string] int {

	contadorVenda := make(map[string] int)
	for _, venda := range vendasMiojo {
		_, vendedorExiste := contadorVenda[venda.Nome]
		if vendedorExiste {
			contadorVenda[venda.Nome] += venda.Valor
		} else {
			contadorVenda[venda.Nome] = venda.Valor
		}

	}

	return contadorVenda

}

func initVenda() []Venda {

	agora := time.Now()
	v1 := Venda{
		Valor : 120,
		Nome: "eu",
		Data: agora,
	}

	v2 := Venda{
		Valor : 10,
		Nome: "eu mesmo",
		Data: time.Date(2021, 01, 01, 0,0,0, 0,time.UTC),
	}

	v3 := Venda{
		Valor : 100,
		Nome: "Irene",
		Data: agora.AddDate(0, -1, 0),
	}

	v4 := Venda{
		Valor : 1000,
		Nome: "Irene",
		Data: agora.AddDate(-1, 0, 0),
	}

	return [] Venda {v1, v2, v3, v4}
		

}


func main() {
	vendas := initVenda()

	fmt.Println(totalVendido(vendas))
}