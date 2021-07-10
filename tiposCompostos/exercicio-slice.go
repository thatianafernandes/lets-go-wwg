//declare duas slices de int e preencha 6 posicoes de cada uma, soma a slices, formando uma 3a slice. Printe as 3
package main
import "fmt"

func main() {
	var slice1 []int
	var slice2 []int

	slice1 = append(slice1, 1)
	slice1 = append(slice1, 2)
	slice1 = append(slice1, 3)
	slice2 = append(slice2, 3)
	slice2 = append(slice2, 2)
	slice2 = append(slice2, 1)

	var soma []int
	soma = append(soma, slice1[0] + slice2[0])
	soma = append(soma, slice1[1] + slice2[1])
	soma = append(soma, slice1[2] + slice2[2])

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(soma)
}