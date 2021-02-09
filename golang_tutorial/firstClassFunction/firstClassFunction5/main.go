package main

import (
	"fmt"
)

func appendStr() func(string) string  {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func main()  {
	a := 5
	func() {
		fmt.Println("a =", a)
	}()

	b := appendStr()
	c := appendStr()
	fmt.Println(b("World"))
	fmt.Println(c("Everyone"))

	fmt.Println(b("Gopher"))
	fmt.Println(c("!"))
}