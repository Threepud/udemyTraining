package main

import "fmt"

func main() {
	foo := "Steve"

	switch foo {
	case "Bob":
		fmt.Println("Hi Bob")
	case "Steve":
		fmt.Println("Hi Steve")
	default:
		fmt.Println("Hi Stranger")
	}
}