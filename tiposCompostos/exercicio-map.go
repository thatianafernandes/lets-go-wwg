//crie um mapa chamado ano onde as chaves são os numeros dos meses e o valor é o nome do mes
//print os mese janeiro e dezembro, printe o tamanho do mapa

package main
import "fmt"

func main() {
	mapa := make(map[string]int)
	mapa["janeiro"] = 1
	mapa["fevereiro"] = 2
	mapa["marco"] = 4
	mapa["abril"] = 5
	mapa["junho"] = 6
	mapa["julho"] = 7
	mapa["agosto"] = 8
	mapa["setembro"] = 9
	mapa["outubro"] = 10
	mapa["novembro"] = 11	
	mapa["dezembro"] = 12

	fmt.Println(mapa)
	fmt.Println(mapa["janeiro"])
	fmt.Println(mapa["dezembro"])
	fmt.Println(len(mapa))

}

