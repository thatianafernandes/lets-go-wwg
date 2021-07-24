package main

import "fmt"

func multiploDe(multiplo int, number int) bool {

    if ((number % multiplo) == 0) {
        fmt.Println("multiplo de", multiplo)
        return true
    }
    return false
    
}
func multiplo(number int) {
    isMultiplo := false
    isMultiplo =  multiploDe(2, number) || isMultiplo
    isMultiplo = multiploDe(3, number) || isMultiplo

    if (!isMultiplo) {
		fmt.Println("nao eh multiplo de 2 ou 3")
	}

}

func main() {
    var number int

    fmt.Scan(&number)
    
	multiplo(number)

}