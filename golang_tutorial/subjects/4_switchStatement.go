package subjects

import (
	"fmt"
	"math/rand"
)

func number() int {
	num2 := 15*5
	return num2
}

func PrintSwitchStatement(print string)  {
	if print == "yes" {
		finger := 4
		fmt.Printf("Finger %d is ", finger)
		// "switch finger := 4;finger {" is correct too
		switch finger {
		case 1:
			fmt.Println("Thumb")
		case 2:
			fmt.Println("Index")
		case 3:
			fmt.Println("Middle")
		case 4:
			fmt.Println("Ring")
		case 5:
			fmt.Println("Pinky")
		default:
			fmt.Println("[incorrect finger number]")
		}
		fmt.Println("")

		letter := "i"
		fmt.Printf("Letter %s is a", letter)
		switch letter {
		case "a", "e", "i", "o", "u":
			fmt.Println("vowel")
		default:
			fmt.Println("not a vowel")
		}
		fmt.Println("")

		num := 75
		switch {
		case num >= 0 && num <= 50:
			fmt.Printf("%d is greater than 0 and less than 50\n", num)
		case num >= 51 && num <= 100:
			fmt.Printf("%d is greather than 51 and less than 100\n", num)
		case num >= 101:
			fmt.Printf("%d is greather than 100\n", num)
		}
		fmt.Println("")

		switch num2 := number();{
		case num2 < 50:
			fmt.Printf("%d is lesser than 50\n", num2)
			fallthrough
		case num2 < 100:
			fmt.Printf("%d is lesser than 100\n", num2)
			fallthrough
		case num < 200:
			fmt.Printf("%d is lesser than 200\n",num)
		}
		fmt.Println("")

		// Will print the second case even if it's false
		switch num3 := 25; {
		case num3 < 50:
			fmt.Printf("%d is lesser than 50\n", num3)
			fallthrough
		case num > 100:
			fmt.Printf("%d is greater than 100\n", num3)
		}
		fmt.Println("")

		switch num4 := -5; {
		case num4 < 50:
			if num4 < 0 {
				break
			}
			fmt.Print("%d is lesser than 50\n", num4)
			fallthrough
		case num4 < 100:
			fmt.Printf("%d is lesser than 100\n", num4)
			fallthrough
		case num4 < 200:
			fmt.Printf("%d is lesser than 200\n", num4)
		}
		fmt.Println("")

	randLoop:
		for {
			switch i := rand.Intn(100); {
			case i % 2 == 0:
				fmt.Printf("Generated even number %d", i)
				break randLoop
			}
		}
	} else {
		// Do nothing
	}
}