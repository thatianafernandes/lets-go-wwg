package main

import "fmt"

func olaMundo() {
	fmt.Println("Hello World")
}

func olaNome(nome string) {
	fmt.Printf("Hello, %v!\n", nome)
}


func cumprimentoPorHora(hora int) string {

	switch  {

		case (hora >= 0 && hora < 6): return "boa madrugada,"
		case hora >= 6 && hora < 12: return "bom dia,"
		case hora >= 12 && hora < 18: return "boa tarde,"
		case hora >= 18 && hora < 24: return "boa noite,"	
		default : return "ola,"
	}

}

func cumprimentoPorHoraDoDia(nome string, hora int) {
	cumprimentoDaHora := cumprimentoPorHora(hora)

	fmt.Println(cumprimentoDaHora, nome)
}

func main() {
	olaMundo()
	olaNome("eu")
	cumprimentoPorHoraDoDia("eu", 01)
	cumprimentoPorHoraDoDia("eu", 07)
	cumprimentoPorHoraDoDia("eu", 13)
	cumprimentoPorHoraDoDia("eu", 19)
}