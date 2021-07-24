package main

import (
	"fmt"
	"strconv"
)

// CPF simula o documento de identificação de pessoa física brasileira
type CPF string

func main() {
	cpf := CPF("11144477735")
	if cpf.EhValido() {
		fmt.Println("CPF válido")
	} else {
		fmt.Println("CPF inválido")
	}
}

func calculaDigito(soma int) int {
	resto := soma % 11;
	if (resto < 2) {
		return 0;
	}
	return 11 - resto;

}

func converteParaDigito(x byte) int{
	return int(x - '0')
}


func (c CPF) ValidoAlgoritmoCPF() bool {
	somaDigito1 := 0
	somaDigito2 := 0
	for i := 0; i < len(c)-2; i++ {
		digito := converteParaDigito(c[i])
		pesoDigito1 := 10-i;
		pesoDigito2 := 11-i;
		somaDigito1 += pesoDigito1 * digito
		somaDigito2 += pesoDigito2 * digito
	}
	primeiroDigito := calculaDigito(somaDigito1)
	somaDigito2 += 2 * primeiroDigito
	segundoDigito := calculaDigito(somaDigito2)


	return (converteParaDigito(c[len(c)-2]) == primeiroDigito) && 
	(converteParaDigito(c[len(c)-1]) == segundoDigito)

}

// EhValido retorna se um CPF é valido ou não
func (c CPF) EhValido() bool {
	_, err := strconv.Atoi(string(c))
	if err == nil && len(c) == 11 {
		return c.ValidoAlgoritmoCPF()
	}
	return false
}