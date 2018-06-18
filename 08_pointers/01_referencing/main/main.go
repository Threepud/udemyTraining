package main

import "fmt"

func main() {
	a := 42

	fmt.Println("a - ", a)
	fmt.Println("a address - ", &a)

	b := &a

	fmt.Println("b - ", b)
	fmt.Println("b - ", *b)

	*b = 100

	fmt.Println("a - ", a)
}
