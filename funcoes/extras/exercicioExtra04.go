package main

import "fmt"


func add(texto string,  value int) string {
	var novoTexto string
	for i := 0; i < len(texto); i++ {
		novoTexto += string(texto[i]+byte(value))
	}
	return novoTexto
}

func add3(texto string) string {
	return add(texto, 3)
}


func minus3(texto string) string {
	return add(texto, -3)
}

func main() {
	text := "oi"
	fmt.Println(add3(text))
	fmt.Println(minus3(text))
	fmt.Println(minus3(add3(text)))
}