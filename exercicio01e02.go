package main

import "fmt"

func main() {
	var myVar1, myVar2, myVar3 int

	fmt.Printf("%v\n", myVar1)
	fmt.Printf("%v\n", myVar2)
	fmt.Printf("%v\n", myVar3)

	a := 230
	b := 27

	fmt.Printf("%v - %v = %v\n", a, b, a-b)
	fmt.Printf("%v + %v = %v\n", a, b, a+b)

	soma := a + b
	subtracao := a - b;
	
	fmt.Printf("%v\n", soma)
	fmt.Printf("%v\n", subtracao)



}