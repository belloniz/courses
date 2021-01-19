package subjects

import "fmt"

func PrintLoops(print string)  {
	if print == "yes" {
		for i := 1; i <= 10; i++ {
			fmt.Printf(" %d",i)
		}
		fmt.Print("\n")

		for i := 1; i <= 10; i++ {
			if i > 5 {
				break
			}
			fmt.Printf(" %d", i)
		}
		fmt.Print("\n")

		for i := 1; i <= 10; i++ {
			if i % 2 == 0 {
				continue
			}
			fmt.Printf(" %d", i)
		}
		fmt.Print("\n")

		n := 5
		for i := 0; i < n; i++ {
			for j := 0; j <= i; j++ {
				fmt.Print(" *")
			}
			fmt.Println()
		}

	outer:
		for i := 0; i < 3; i++ {
			for j := 1; j < 4; j++ {
				fmt.Printf(" i = %d , j = %d\n", i, j)
				if i == j {
					break outer
				}
			}
		}
	} else {
		// Do nothing
	}
}