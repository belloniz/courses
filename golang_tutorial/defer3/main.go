package main

import "fmt"

func main()  {
	name := "John"
	fmt.Printf("Original String: %s\n", string(name))
	fmt.Printf("Reversed String: ")
	for _, v := range name {
		defer fmt.Printf("%c", v)
	}
}