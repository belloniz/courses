package subjects

import "fmt"

func subtractOne(numbers []int)  {
	for i := range numbers {
		numbers[i] -=2
	}
}

func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	fmt.Println(neededCountries)
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries)
	return countriesCpy
}

func PrintSlices(print string) {
	if print == "yes" {
		a := [5]int{76, 77, 78, 79, 80}
		var b []int = a[1:4]
		fmt.Println(b)
		fmt.Println("")

		c := []int{6, 7, 8}
		fmt.Println(c)
		fmt.Println("")

		darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
		dslice := darr[2:5]
		fmt.Println("array before", darr)
		for i := range dslice {
			dslice[i]++
		}
		fmt.Println("array after", darr)
		fmt.Println("")

		numa := [3]int{78, 79, 80}
		nums1 := numa[:]
		nums2 := numa[:]
		fmt.Println("array before change 1", numa)
		nums1[0] = 100
		fmt.Println("array after modification to slice 1", numa)
		nums2[1] = 101
		fmt.Println("array after modification to slice 2", numa)
		fmt.Println()

		fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
		fruitslice := fruitarray[1:3]
		fmt.Printf("lenght of slice %d capacity %d", len(fruitslice), cap(fruitslice))
		fmt.Println()
		fruitslice = fruitslice[:cap(fruitslice)]
		fmt.Println("After re-slicing lenght is", len(fruitslice), "and capacity is", cap(fruitslice))
		fmt.Println()

		i := make([]int, 5, 5)
		fmt.Println(i)
		fmt.Println()

		cars := []string{"Ferrari", "Honda", "Ford"}
		fmt.Println("cars:", cars, "has old lenght", len(cars), "and capacity", cap(cars))
		cars = append(cars, "Toyota")
		fmt.Println("cars:", cars, "has new lenght", len(cars), "and new capacity", cap(cars))
		fmt.Println()

		var names []string
		if names == nil {
			fmt.Println("slice is nil going to append")
			names = append(names, "John", "Sebastian", "Vinay")
			fmt.Println("names contents:",names)
		}
		fmt.Println()

		veggies := []string{"potatoes", "tomatoes", "brinjal"}
		fruits := []string{"oranges", "apple"}
		food := append(veggies, fruits...)
		fmt.Println("food:", food)
		fmt.Println()

		nos := []int{8, 7, 6}
		fmt.Println("slice before function call", nos)
		subtractOne(nos)
		fmt.Println("slice after function call", nos)
		fmt.Println()

		pls := [][]string {
			{"C","C++"},
			{"Javascript"},
			{"Go","Rust"},
		}

		for _, v1 := range pls {
			for _, v2 := range v1 {
				fmt.Printf("%s ", v2)
			}
			fmt.Printf("\n")
		}
		fmt.Println()

		countriesNeeded := countries()
		fmt.Println(countriesNeeded)
		
	} else {
		// Do nothing
	}
}