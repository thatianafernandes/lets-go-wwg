//CALCULE O VALOR TOTAL DE UMA COMPRA QUE TEM 3 ITENS REPRESENTANDO O PREÇO DE TODOS ELES EM FLOAT64
//TODOS OS ITENS DESSA COMPRA PRECISAM SER COMPRADOS EM MAIS DE UMA UNIDADE

package main

import "fmt"

func leProduto() (int, float64) {
	var preco float64
	var quantidade int

	fmt.Scanln(&quantidade, &preco)

	return quantidade, preco

}

func main() {
	var preco1, preco2, preco3 float64

	var quantidade1, quantidade2, quantidade3 int

	quantidade1, preco1 = leProduto()
	quantidade2, preco2 = leProduto()
	quantidade3, preco3 = leProduto()

	var soma float64

	soma = float64(quantidade1) * preco1 + float64(quantidade2) * preco2 + float64(quantidade3) * preco3


	fmt.Println(soma)


	

}