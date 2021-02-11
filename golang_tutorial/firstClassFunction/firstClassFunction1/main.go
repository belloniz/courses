package main

import (
	"fmt"
)

func main()  {
	a := func() {
		fmt.Println("hello world first class function")
	}
	a()
	fmt.Printf("%T", a)
	fmt.Println()

	func() {
		fmt.Println("hello world first class function")
	}()

	func(n string) {
		fmt.Println("Welcome", n)
	}("Gophers")
}