//declare duas variaves do tipo string, uma vai receber seu nome e outra sua cor preferida
//printe na tela uma frase usando os dois valres usando o fmt.Print e o especificador de formato %s
//usando uma linha por variável, demonstre o valor e o tipo de cada uma delas

package main

import "fmt"

func main() {
	var nome, cor string

	nome = "thatiana"

	cor = "azul"

	fmt.Print(nome, " gosta da cor ",  cor, "\n")
	fmt.Printf("%s gosta da cor %s\n", nome, cor)
	fmt.Println(nome,"gosta da cor",  cor, "\n")

	fmt.Printf("%s -- %T\n", nome, nome)
	fmt.Printf("%s -- %T\n", cor, cor)

}