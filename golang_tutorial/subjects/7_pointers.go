package subjects

import "fmt"

func changeValueOfPointer(val *int) {
	*val = 55
}

func helloPointer() *int {
	i := 5
	return &i
}

func modify(sls []int)  {
	sls[0] = 90
}

func PrintPointers(print string)  {
	if print == "yes" {
		fmt.Println("Declaring pointers")
		b := 255
		var a *int = &b
		fmt.Printf("Type of a is %T\n",a)
		fmt.Println("addres of b is", a)
		fmt.Println(a)
		fmt.Println()

		fmt.Println("Zero value of a pointer")
		a2 := 25
		var b2 *int
		if b2 == nil {
			fmt.Println("b2 is", b2)
			b2 = &a2
			fmt.Println("b2 after initializated is",b2)
		}
		fmt.Println()

		fmt.Println("Creating pointers using the new function")
		size := new(int)
		fmt.Printf("Size value is %d, type is %T, addres is %v\n", *size, size, size)
		*size = 85
		fmt.Println("New size value is", *size)
		fmt.Println()

		fmt.Println("Dereferencing a pointer")
		b3 := 255
		a3 := &b3
		fmt.Println("addres of b3 is", a3)
		fmt.Println("value of b3 is", *a3)
		*a3++
		fmt.Println("new value of b3 is", *a3)
		fmt.Println(b3)
		fmt.Println()

		fmt.Println("Passing pointer to a function")
		a4 := 58
		fmt.Println("Value of a4 before function call is", a4)
		b4 := &a4
		changeValueOfPointer(b4)
		fmt.Println("Value of a4 after function call is", a4)
		fmt.Println()

		fmt.Println("Returning pointer from a function")
		d := helloPointer()
		fmt.Println("Value of d", *d)
		fmt.Println()

		fmt.Println("Do not pass a pointer to an array as a argument to a function. Use slice instead.")
		a5 := [3]int{89, 90, 91}
		modify(a5[:])
		fmt.Println(a5)
		fmt.Println()

		fmt.Println("Go does not support pointer arithmetic")
		// b := [...]int{109, 110, 111}
		// p := &b
		// p++
		// This will throw an error
	} else {
		// Do nothing
	}
}
