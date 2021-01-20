package subjects

import (
	"fmt"
)

func hello(a int, b ...int)  {
}

func find(num int, nums ...int)  {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Println()
}

func change(s...string)  {
	s[0] = "Go"
}

func change2(s...string){
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
}

func PrintVariadicFunctions(print string)  {
	if print == "yes" {
		hello(1)
		hello(1,2)
		hello(5, 6, 7, 8, 9)
		fmt.Println()

		find(89, 89, 90, 95)
		find(45, 56, 67, 45, 90, 109)
		find(78, 38, 56, 98)
		find(87)
		nums := []int{89, 90, 95}
		find(89, nums...)
		fmt.Println()

		welcome := []string{"hello","world"}
		change(welcome...)
		fmt.Println(welcome)
		fmt.Println()

		welcome2 := []string{"hello","world"}
		change2(welcome2...)
		fmt.Println(welcome2)

	} else {
		// Do nothing
	}
}
