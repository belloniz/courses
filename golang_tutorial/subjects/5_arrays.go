package subjects

import (
	"fmt"
)

func changeLocal(num [5]int)  {
	num[0] = 55
	fmt.Println("inside function ",num)
}

func printarray(a9 [3][2]string)  {
	for _, v1 := range a9 {
		for _, v2 := range v1 {
			fmt.Printf("%s ",v2)
		}
		fmt.Printf("\n")
	}
}

func PrintArrays(print string)  {
	if print == "yes" {
		var a [3]int
		a[0] = 12
		a[1] = 78
		a[2] = 50
		fmt.Println(a)
		fmt.Println("")

		a2 := [3]int{12,78,50}
		fmt.Println(a2)
		fmt.Println("")

		a3 := [3]int{12}
		fmt.Println(a3)
		fmt.Println("")

		a4 :=[]int{12, 78, 50}
		fmt.Println(a4)
		fmt.Println("")

		a5 := [...]string{"USA", "China", "India", "Germany", "France"}
		b := a5
		b[0] = "Singapore"
		fmt.Println("a is", a5)
		fmt.Println("b is", b)
		fmt.Println("")

		num := [...]int{5, 6, 7, 8, 8}
		fmt.Println("before passing to function ", num)
		changeLocal(num)
		fmt.Println("after passing to function ", num)
		fmt.Println("")

		a6 := [...]float64{67.7, 89.8, 21, 78}
		fmt.Println("length of a6 is", len(a6))
		fmt.Println("")

		a7 := [...]float64{67.7, 89.8, 21, 78}
		for i := 0; i < len(a7); i++ {
			fmt.Printf("%d th element of a is %.2f\n", i, a7[i])
		}
		fmt.Println("")

		a8 := [...]float64{67.7, 89.8, 21, 78}
		sum := float64(0)
		for i, v := range a8 {
			fmt.Printf("%d the element of a is %2.f\n", i, v)
			sum += v
		}
		fmt.Println("\nsum of all elements of a", sum)
		fmt.Println("")

		a9 := [3][2]string{
			{"lion","tiger"},
			{"cat","dog"},
			{"pigeon","peacock"},
		}

		printarray(a9)

		var b2 [3][2]string
		b2[0][0] = "apple"
		b2[0][1] = "samsung"
		b2[1][0] = "microsoft"
		b2[1][1] = "google"
		b2[2][0] = "AT&T"
		b2[2][1] = "T-Mobile"
		fmt.Printf("\n")
		printarray(b2)
	} else {
		// Do nothing
	}
}