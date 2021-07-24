package main

import "fmt"

type MySlice [] int

func (slice MySlice) soma() int {
	soma := 0
	for i := 1; i <len(slice); i++ {
		soma += slice[i]
	}
	return soma
}

func (slice MySlice) media() float64 {	
	return float64(slice.soma())/float64(len(slice))
}

func main() {
	s := MySlice {1, 2, 2, 3, 10}

	fmt.Println(s.soma())
	fmt.Println(s.media())	


}