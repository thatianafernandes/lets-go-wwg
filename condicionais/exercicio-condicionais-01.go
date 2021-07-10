package main

import (
	"fmt"
	"time"
)

func scanVar(nome string,  min int,  max int) int {
    var timeFrame int
    fmt.Println(nome + ":")
    n, err := fmt.Scanln(&timeFrame)
    if (err != nil) || (timeFrame < min || timeFrame > max) || (n != 1) {
        panic(err)
    }

    return timeFrame
}

func scanBorn(todayYear int) (int, int, int) {	
	year := scanVar("Ano", 0, todayYear)    
    month := scanVar("Mes", 1, 12) 
	day := scanVar("Dia", 1, 31) //preguica de ver o dia...
	
	return year, month, day
}

func calcVotingAge(today time.Time,  year int,  month int, day int) bool {	
	const votingAge = 16;
	todayYear := today.Year()
	age := todayYear - year

	monthInt := int(today.Month())

	if (month > monthInt) || (month == monthInt && day > today.Day()) {
		age = age - 1
	}

	return age >= votingAge;

}

func main() {
    fmt.Println("Coloque o dia, mes e ano em que nasceu.")
	today := time.Now()
	year, month, day := scanBorn(today.Year())

	if(!calcVotingAge(today, year, month, day)) {
		fmt.Println("Nao apto a votar")
	} else {
		fmt.Println("Apto a votar")
	}
	
}