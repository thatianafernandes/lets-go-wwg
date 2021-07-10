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

func calcAge(today time.Time,  year int,  month int, day int) string {	
	todayYear := today.Year()
	if year == todayYear {
		return "menos de 1 ano"
	}
	
	age := todayYear - year

	monthInt := int(today.Month())

	if (month > monthInt) || (month == monthInt && day > today.Day()) {
		age = age - 1
	}

	return fmt.Sprint(age , " anos")

}

func main() {
    fmt.Println("Coloque o dia, mes e ano em que nasceu.")
	today := time.Now()
	year, month, day := scanBorn(today.Year())

	fmt.Println(calcAge(today, year, month, day))
}