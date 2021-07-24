package main

import (
	"fmt" 
	"strings"
	"os"
	"bufio"
	"unicode"
)

func leiaLinha() string {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	return sc.Text()
}



func contaLetras(texto string) map[rune] int {
	contador := make(map[rune] int)

	for _,letra := range texto {		
		quantidade, letraJaExiste := contador[letra]
		if letraJaExiste {
			contador[letra] = quantidade + 1
		} else {
			contador[letra] = 1
		}		
	}

	return contador


}

func contaLetrasIgnorandoCase(texto string) map[rune] int {
	texto = strings.ToLower(texto);
	return contaLetras(texto);

}

func imprimeLetraSeExiste(contadorLetras map[rune] int, letra rune) {
	quantidade, letraExiste := contadorLetras[letra]
	if letraExiste {
		fmt.Printf("%c - %d\n", letra, quantidade)
	}	
}

func main() {

	texto := leiaLinha()

	//case sensitive
	fmt.Println("Case sensitve:")
	contadorLetras := contaLetras(texto)

	for letra := 'a'; letra <= 'z'; letra++ {
		imprimeLetraSeExiste(contadorLetras, letra)	
		letraMaiuscula := unicode.ToUpper(letra)
		imprimeLetraSeExiste(contadorLetras, letraMaiuscula)	
	}

	//case insensitive
	fmt.Println("Case insensitve:")
	contadorLetras = contaLetrasIgnorandoCase(texto)

	for letra := 'a'; letra <= 'z'; letra++ {
		imprimeLetraSeExiste(contadorLetras, letra)	
	}	

}