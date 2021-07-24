package main

import "fmt"

func main() {

	for linha := 1; linha <= 10; linha++ {
		numeroNaLinha := 1;
		for numeroNaLinha <= linha {
			fmt.Print(numeroNaLinha)
			numeroNaLinha++
		} 
		fmt.Print("\n")
	}

}