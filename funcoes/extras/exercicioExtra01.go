package main

import "fmt"

func modificaLista(listaInteiros * []int) (int, int) {
	somaAntes := 0;
	somaDepois := 0

	for i := 0; i < len(*listaInteiros); i++ {		
		somaAntes += (*listaInteiros)[i]
		if ((*listaInteiros)[i] % 2 == 0) {
			(*listaInteiros)[i] = (*listaInteiros)[i] / 2
		} else  {
			(*listaInteiros)[i] = (*listaInteiros)[i] *  2
		}
		somaDepois += (*listaInteiros)[i]

	}



	return somaAntes, somaDepois
}

func main() {
	listaInteiros := [] int {1, -2, -3, 4, 7, 10}

	antes, depois := modificaLista(&listaInteiros)

	fmt.Println(listaInteiros, antes, depois)


}