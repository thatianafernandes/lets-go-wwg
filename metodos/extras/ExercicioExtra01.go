package main

import (
	"fmt"
	"errors"
)

type Stack []int

func (s * Stack) push(e int) {
	*s = append(*s, e)
}

func (s * Stack) pop() (int, error) {

	if (len(*s) == 0) {
		return -1, errors.New("cannot pop from empty stack")
	}

	elem := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return elem, nil

}

func main() {
	var myStack Stack

	myStack.push(1)
	myStack.push(2)
	myStack.push(3)
	myStack.push(4)

	fmt.Println(myStack)

	myStack.pop()
	myStack.pop()
	fmt.Println(myStack)

	myStack.pop()
	myStack.pop()
	_, err := myStack.pop()

	if (err != nil) {
		fmt.Println("Empty stack")
	}

}