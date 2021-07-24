package main

import "fmt"

func initMap() map[string] string {
	mapa := make( map[string] string)
	mapa["Rio de Janeiro"] = "Brasil"
	mapa["Lisboa"] = "Portugal"
	mapa["Tel Aviv"] = "Israel"
	mapa["Bruxelas"] = "Belgica"
	mapa["Porto"] = "Portugal"
	mapa["Sao Paulo"] = "Brasil"
	mapa["Buenos Aires"] = "Argentina"
	mapa["Bogota"] = "Colombia"
	mapa["Roma"] = "Italia"
	mapa["Berlim"] = "Alemanha"
	mapa["Salta"] = "Argentina"

	return mapa
}

func countPaises(mapa map[string] string)  map[string] int {
	count := make( map[string] int)
	for _, pais := range mapa {
		quantidade, paisJaExiste := count[pais]
		if paisJaExiste {
			count[pais] = quantidade + 1
		} else {
			count[pais] = 1
		}
	}

	return count
	
}

func main() {

	mapa := initMap()

	count := countPaises(mapa)

	fmt.Println(count)

}