package main

import "fmt"

func main() {

	theAnswer := 42

	fmt.Println("a - ", theAnswer)
	fmt.Println("a Address - ", &theAnswer)
	fmt.Printf("%d\n", &theAnswer)

}
