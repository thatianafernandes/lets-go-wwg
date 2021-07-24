package main

import "fmt"

func changeVars(a int, b int) (int, int) {
	aux := a
	a = b
	b = aux

	return a, b
}


func main() {
	var a, b, c int

	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	if a > b {
		a, b = changeVars(a, b)
	}
	if a > c {
		a, c = changeVars(a, c)
	}
	if b > c {
		b, c = changeVars(b, c)
	}



	fmt.Println(a, b, c)
}