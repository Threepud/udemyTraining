package main

import "fmt"

func main() {
	for i := 5000; i < 5100; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}

	foo := "a"
	fmt.Println(foo)
	fmt.Printf("%T\n", foo)
	fmt.Printf("%T\n", foo[0])
	fmt.Printf("%T\n", rune(foo[0]))
}