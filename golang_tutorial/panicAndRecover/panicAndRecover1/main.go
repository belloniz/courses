package main

import (
	"fmt"
)

func recoverFullName()  {
	if r := recover(); r != nil {
		fmt.Println("recovered from", r)
	}
}

func invalidSliceAccess()  {
	defer recoverFullName()
	n := []int{5, 7, 4}
	fmt.Println(n[4])
	fmt.Println("normally returned from a")
}

func fullName(firstName *string, lastName *string) {
	defer recoverFullName()
	if firstName == nil {
		panic("runtime error: first name cannot be nill")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nill")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from Fullname")
}

func main()  {
	defer fmt.Println("deffered call in main")
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")

	invalidSliceAccess()
	fmt.Println("normally returned from main")
}