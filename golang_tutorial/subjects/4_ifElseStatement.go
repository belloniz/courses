package subjects

import (
	"fmt"
)

func PrintIfElseStatement(print string)  {
	if print == "yes" {
		num := 99
		if num%2 == 0 {
			fmt.Println("The number", num, "is even")
			return
		} else {
			fmt.Println("The number", num, "is odd")
		}

		if num <= 50 {
			fmt.Println(num, "is less than or equal to 50")
		} else if num >= 51 && num <= 100 {
			fmt.Println(num, "is between 51 and 100")
		} else {
			fmt.Println(num, "is greather than 100")
		}

		if num2 := 10; num % 2 == 0 {
			fmt.Println(num2, "is even")
		} else {
			fmt.Println(num, "is odd")
		}

		num3 := 11

		// Go's Phisolophy
		if num3 % 2 == 0 {
			fmt.Println(num, "is even")
			return
		}
		fmt.Println(num, "is odd")
	} else {
		// Do nothing
	}
}