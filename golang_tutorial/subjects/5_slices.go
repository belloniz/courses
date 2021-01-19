package subjects

import (
	"fmt"
)

func PrintSlices() {
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


}