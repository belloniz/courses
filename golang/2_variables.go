package main

// This is cool and important, is the same use of the multiple variable declaration
import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Declaring a single variable")
	var age int // variable declaration
	fmt.Println("My age is", age)
	age = 29
	fmt.Println("My age is", age)
	age = 54
	fmt.Println("My age is", age)
	fmt.Println("")

	fmt.Println("Declaring a variable with an initial value")
	var age2 int = 2 // variable declaration with initial value
	fmt.Println("My age is", age2)
	fmt.Println("")
	
	fmt.Println("Type inference")
	var age3 = 29
	fmt.Println("My age3 is ", age3)
	fmt.Println("")

	fmt.Println("Multiple variable declaration")
	var width, height int = 100, 50 // declaring multiple variables
	fmt.Println("width is", width, "height is",height)

	var width2, height2  = 100, 50 // declaring multiple variables
	fmt.Println("width2 is", width2, "height2 is",height2)

	var width3, height3 int
	fmt.Println("width3 is", width3, "height3 is",height3)
	width3 = 333
	height3 = 33
	fmt.Println("width3 is", width3, "height3 is",height3)

	var ( // declaring multiples variables but with differents types
		nameG = "Gabriel"
		ageG = 23
		heightG int
	)
	fmt.Println("my name is", nameG,"my age is", ageG, "my height is", heightG)
	fmt.Println("")

	fmt.Println("Short hand declaration")
	count := 10
	fmt.Println("Count =",count)

	name, age4 := "Gabriel", 23 // short hand declaration
	fmt.Println("my name is", name,"my age4 is", age4)

	a, b := 20, 30 // declare variables a and b
    fmt.Println("a is", a, "b is", b)
    b, c := 40, 50 // b is already declared but c is new
    fmt.Println("b is", b, "c is", c)
    b, c = 80, 90 // assign new values to already declared variables b and c
	fmt.Println("changed b is", b, "c is", c)
	
	a2, b2 := 145.8, 543.8
	c2 := math.Min(a2,b2) // math is a package and Min is a function in that package
	fmt.Println("Minimum value is", c2)
}