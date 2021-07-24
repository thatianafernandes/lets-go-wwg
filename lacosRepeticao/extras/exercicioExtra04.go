package main

import "fmt"

func main() {

	IsNumeroPar := false
	for !IsNumeroPar {
		var x int
		fmt.Scanln(&x)
		if x % 2 == 0 {
			fmt.Println("Agora sim podemos dividir igualmente entre mim e você!")
			IsNumeroPar = true
		} 

	}

}