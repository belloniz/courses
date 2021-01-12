package main

import (
	"fmt"
	"math"
)

func main(){
	const a = 50
	fmt.Println(a)
	fmt.Println("")

	fmt.Println("Declaring a group of constants")
	const (
		name = "John"
		age = 50
		country = "Canada"
	)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(country)

	var a2 = math.Sqrt(4)   //allowed
	//const b2 = math.Sqrt(4) //not allowed
	fmt.Println(a2)
	fmt.Println("")

	fmt.Println("String Constants, Typed and Untyped Constants")
	const typedHello string = "Hello World"
	fmt.Println(typedHello)
	var defaultName = "Sam" // allowed
	type myString string // allowed
	var customName myString = "Sam" // allowed
	fmt.Println(defaultName, customName)
	fmt.Println("")

	fmt.Println("Boolean Constants")
	const trueConst = true
	type myBool bool
	var defaultBool = trueConst
	var customBool myBool = trueConst
	fmt.Println(defaultBool,customBool)
	fmt.Println("")

	fmt.Println("Numeric Constants")
	const a3 = 5
	var intVar int = a3
	var int32Var int32 = a3
	var float64Var float64 = a3
	var complex64Var complex64 = a3
	fmt.Printf("intVar", intVar, "\nintvar32", int32Var, "\nfloat64", float64Var, "\ncomplex64", complex64Var)
	fmt.Println("")

	var i = 5
	var f = 5.6
	var c = 5 + 6i
	fmt.Printf("i's type is %T, f's type is %T, c's type is %T", i, f, c)
	fmt.Println("\n")

	const a4 = 5
	var intVar_2 int = a4
	var int32Var_2 int32 = a4
	var float64Var_2 float64 = a4
	var complex64Var_2 complex64 = a4
	fmt.Println("intVar", intVar_2, "\nint32Var", int32Var_2, "float64Var", float64Var_2, "complex64Var", complex64Var_2) 
	fmt.Println("")

	fmt.Println("Numeric Expressions")
	var a5 = 5.9 / 8 
	fmt.Printf("a's type if %T and value is %v", a5, a5)

}