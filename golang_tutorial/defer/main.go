package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName string
}

func finished()  {
	fmt.Printf("Finished finding largest.")
}

func largest(nums []int) {
	defer finished()
	fmt.Println("Started finding larges")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Largest number in", nums, "is", max)
}

func (p person) fullName()  {
	fmt.Println(p.firstName + " " + p.lastName)
}

func main()  {
	nums := []int{78, 109, 2, 563, 300}
	largest(nums)
	fmt.Println()
	fmt.Println()

	p := person{"John", "Wick",}
	defer p.fullName()
	fmt.Printf("Welcome ")
}