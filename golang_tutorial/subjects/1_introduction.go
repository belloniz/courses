package subjects

import "fmt"

func PrintIntroduction(print string){
	if print == "yes" {
		fmt.Println("Hello World")
		fmt.Println()
	} else {
		// Do nothing
	}
}