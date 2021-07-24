package main

import "fmt"

func ex01() {

	for i := 13; i <=27; i++ {
		fmt.Println(i)
	}
}

func ex02() {
	h := 1
	for h <= 24  {
		m := 0
		for m < 60 {
			s := 1
			for s < 60 {
				fmt.Printf("%v:%2v:%2v\n", h, m, s)
				s++
			}
			
			m++
		}
		
		h++
	}
}

func main() {
	ex01();
	ex02();
}