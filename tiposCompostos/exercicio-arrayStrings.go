// array de strings o valor de cada elemento deve ser o numero do indice escrito por extenso. Printe na tela o tipo do seu array e os valores dos seus elementos
package main
import "fmt"

func main() {
	var numeros [6]string
	numeros[0] = "zero"
	numeros[1] = "um"
	numeros[2] = "dois"
	numeros[3] = "três"
	numeros[4] = "quatro"
	numeros[5] = "cinco"

	fmt.Printf("o tipo é: %T\n", numeros)
	fmt.Println(numeros)
}