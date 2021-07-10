package main
import "fmt"

func main() {
	zodiaco := make([]string, 12)
	zodiaco[0] = "Aries"
	zodiaco[1] = "Touro"
	zodiaco[2] = "Gemeos"
	zodiaco[3] = "Cancer"
	zodiaco[4] = "Leao"
	zodiaco[5] = "Virgem"
	zodiaco[6] = "Libra"
	zodiaco[7] = "Escorpiao"
	zodiaco[8] = "Sagitario"
	zodiaco[9] = "Capricornio"
	zodiaco[10] = "Aquario"
	zodiaco[11] = "Peixes"


	fmt.Println("tamanho do slice: ", len(zodiaco))
	fmt.Println("capacidade do slice: ", cap(zodiaco))
	fmt.Println("elementos do slice: ", zodiaco)
}