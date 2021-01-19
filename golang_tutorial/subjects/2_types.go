package subjects

import (
	"fmt"
	"unsafe"
)

func PrintTypes(print string)  {
	if print == "yes"{
		fmt.Println("bool")
		a := true
		b := false
		fmt.Println("a:", a, "b:", b)

		c := a && b
		fmt.Println("c:", c)

		d := a || b
		fmt.Println("d:", d)
		fmt.Println("")

		fmt.Println("Signed integers")
		//int8: represents 8 bit signed integers
		//size: 8 bits
		//range: -128 to 127

		//int16: represents 16 bit signed integers
		//size: 16 bits
		//range: -32768 to 32767

		//int32: represents 32 bit signed integers
		//size: 32 bits
		//range: -2147483648 to 2147483647

		//int64: represents 64 bit signed integers
		//size: 64 bits
		//range: -9223372036854775808 to 9223372036854775807

		var a2 = 89
		b2 := 95
		fmt.Println("value of a2 is", a2, "and b2 is", b2)
		fmt.Printf("type of a2 is %T, size of a2 is %d", a2, unsafe.Sizeof(a2)) // type and size of a
		fmt.Printf("\ntype of b2 is %T, size of b2 is %d", b2, unsafe.Sizeof(b2)) // type and size of b
		fmt.Println("")

		fmt.Println("Unsigned integers")
		//uint8: represents 8 bit unsigned integers
		//size: 8 bits
		//range: 0 to 255

		//uint16: represents 16 bit unsigned integers
		//size: 16 bits
		//range: 0 to 65535

		//uint32: represents 32 bit unsigned integers
		//size: 32 bits
		//range: 0 to 4294967295

		//uint64: represents 64 bit unsigned integers
		//size: 64 bits
		//range: 0 to 18446744073709551615

		//uint : represents 32 or 64 bit unsigned integers depending on the underlying platform.
		//size : 32 bits in 32 bit systems and 64 bits in 64 bit systems.
		//range : 0 to 4294967295 in 32 bit systems and 0 to 18446744073709551615 in 64 bit systems
		fmt.Println("")

		fmt.Println("Floating point types")
		//float32: 32 bit floating point numbers
		//float64: 64 bit floating point numbers

		a3, b3 := 5.67, 8.97
		fmt.Printf("type of a3 %T b %T\n", a3, b3)
		sum := a3 + b3
		diff := a3 - b3
		fmt.Println("sum", sum, "diff", diff)
		fmt.Println("")

		fmt.Println("Complex types")
		c1 := complex(5,7)
		c2 := 8 + 27i
		cadd := c1 + c2
		fmt.Println("sum:",cadd)
		cmul := c1 * c2
		fmt.Println("product:",cmul)
		fmt.Println("")

		fmt.Println("Other numeric types")
		//byte is an alias of uint8
		//rune is an alias of int32
		fmt.Println("")

		fmt.Println("String type")
		first := "Naveen"
		last := "Ramanthan"
		name := first + " " + last
		fmt.Println("My name is", name)
		fmt.Println("")

		fmt.Println("Type Conversion")
		i := 55		// int
		j := 67.8	// float
		sum2 := i + int(j) // j is converted to int
		fmt.Println("sum:", sum2)

		i2 := 10
		var j2 float64 = float64(i2)
		fmt.Println("j2:",j2)
	} else {
		// Do nothing
	}
}