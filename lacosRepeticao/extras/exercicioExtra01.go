package main

import "fmt"

func compras() {

	var compras = []string {"vinho", "cerveja", "vodka", "cachaca", "run"}
	compras = append(compras, "tequila")

	for i := 0; i < len(compras); i++  {
		fmt.Println(i, compras[i])
	}	

}

func main() {
	compras()
}