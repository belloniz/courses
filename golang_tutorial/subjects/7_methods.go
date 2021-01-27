package subjects

import (
	"fmt"
	"math"
)

type Employee struct {
	name string
	salary int
	currency string
}

type Employee2 struct {
	name string
	age int
}

type Rectangle struct {
	length int
	width int
}

type Circle struct {
	radius float64
}

type address struct {
	city string
	state string
}

type person struct {
	firstName string
	lastName string
	address
}

type myInt int

func (e Employee)  displaySalary(){
	fmt.Printf("Salary of %s id %s%d", e.name, e.currency, e.salary)
}

func displaySalary2(e Employee){
	fmt.Printf("Salary of %s id %s%d", e.name, e.currency, e.salary)
}

func (r Rectangle) Area() int {
	return r.width * r.length
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (e Employee2) changeName(newName string) {
	e.name = newName
}

func (e *Employee2) changeAge(newAge int) {
	e.age = newAge
}

func (a address) fullAddres() {
	fmt.Printf("Full addres: %s, %s.\n", a.city, a.state)
}

func area(r Rectangle)  {
	fmt.Printf("Area Function result: %d\n", (r.length * r.width))
}

func (r Rectangle) area()  {
	fmt.Printf("Area method result: %d\n", (r.length * r.width))
}

func perimeter(r *Rectangle)  {
	fmt.Println("perimeter function output:", 2*(r.length+r.width))
}

func (r *Rectangle) perimeter()  {
	fmt.Println("perimeter function output:", 2*(r.length+r.width))
}

func (a myInt) add(b myInt) myInt {
	return a + b
}

func PrintMethods(print string)  {
	if print == "yes" {
		emp1 := Employee {
			name: "Sam Adoof",
			salary: 5000,
			currency: "$",
		}
		emp1.displaySalary()
		fmt.Println()
		displaySalary2(emp1)
		fmt.Println()
		fmt.Println()

		r := Rectangle{10,5}
		fmt.Println("Area of rectangle is", r.Area())

		c := Circle{12}
		fmt.Println("Area of circle is", c.Area())
		fmt.Println()

		e := Employee2{
			name: "Mark Andrew",
			age:  50,
		}
		fmt.Printf("Employee name before change: %s", e.name)
		e.changeName("Michael Andrew")
		fmt.Printf("\nEmployee name after change: %s", e.name)

		fmt.Printf("\n\nEmployee age before change: %d", e.age)
		(&e).changeAge(51)
		fmt.Printf("\nEmployee age after change: %d", e.age)
		fmt.Println()

		fmt.Println("When to use pointer receiver and when to use value receiver\n")
		//Generally, pointer receivers can be used when changes made to the receiver inside the method should be visible to the caller.
		//
		//Pointers receivers can also be used in places where it's expensive to copy a data structure. Consider a struct that has many fields.
		//Using this struct as a value receiver in a method will need the entire struct to be copied which will be expensive. In this case,
		//if a pointer receiver is used, the struct will not be copied and only a pointer to it will be used in the method.
		fmt.Println()

		fmt.Println("Methods of anonymous struct fields")
		p := person {
			"Elon",
			"Musk",
			address{
				"Los Angeles",
				"California",
			},
		}
		p.fullAddres()
		fmt.Println()

		fmt.Println("Value receivers in methods vs Value arguments in functions")
		r2 := Rectangle{10,5}
		area(r2)
		r2.area()

		p2 := &r2
		/*
		   compilation error, cannot use p (type *rectangle) as type rectangle
		   in argument to area
		*/
		//area(p)
		p2.area()
		fmt.Println()

		fmt.Println("Pointer receivers in methods vs Pointer arguments in functions")
		r3 := Rectangle{10,5}
		p3 := r3
		p3.perimeter()
		/*
		   cannot use r (type rectangle) as type *rectangle in argument to perimeter
		*/
		//perimeter(r)
		r3.perimeter()
		fmt.Println()

		fmt.Println("Methods with non-struct receivers")
		num1 := myInt(5)
		num2 := myInt(10)

		sum := num1.add(num2)
		fmt.Println("Sum is", sum)
	} else {
		// Do nothing
	}
}