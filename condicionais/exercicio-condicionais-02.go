package main

import (
	"fmt"
)


func main() {
	var number int

	fmt.Println("Escreva um numero")
	n, err := fmt.Scanln(&number)

	if err != nil || n != 1 {
		panic(err)
	}

	if number < 0 {
		fmt.Println("negativo")
	} else {
		fmt.Println("positivo")
	}
	
}