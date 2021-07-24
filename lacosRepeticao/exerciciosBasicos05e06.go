package main

import "fmt"

func compras() {

	var compras = []string {"vinho", "cerveja", "vodka", "cachaca", "run"}
	compras = append(compras, "tequila")

	for indice, item := range compras {
		fmt.Println(indice, item)
	}	

}

func printFatia() {
	var fatia = []string {"zero", "um", "dois", "três", "quatro"}
	for indice, valor := range fatia {
		fmt.Println(indice, valor)
	}		
}

func main() {

	printFatia()
	compras()

}