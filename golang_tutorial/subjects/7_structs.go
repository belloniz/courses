package subjects

import (
	"fmt"
	"golang_tutorial/structs/computer"
)

type Employees struct {
	firstName string
	lastName string
	age int
	salary int
}

type PersonAnon struct {
	string
	int
}

type Address struct {
	city string
	state string
}

type Person struct {
	name string
	age int
	address Address
}

type name struct {
	firstName string
	lastName string
}

type Person2 struct {
	name string
	age int
	Address
}

func PrintStructs(print string)  {
	if print == "yes" {
		fmt.Println("Creating named structs")
		//creating struct specifying field names
		emp1 := Employees{
			firstName: "Sam",
			age: 25,
			salary: 500,
			lastName: "Anderson",
		}
		//creating struct without specifying field names
		emp2 := Employees{"Thomas","Paul",29,800}

		fmt.Println("Employee 1: ", emp1)
		fmt.Println("Employee 2:", emp2)
		fmt.Println()

		fmt.Println("Creating anonymous structs")
		emp3 := struct {
			firstName string
			lastName string
			age int
			salary int
		}{
			firstName: "Andreah",
			lastName:  "Nikola",
			age:       31,
			salary:    5000,
		}
		fmt.Println("Employee 3:", emp3)
		fmt.Println()

		fmt.Println("Accessing individual fields of a struct")
		fmt.Println("Employee 1")
		fmt.Println("Fist name:", emp1.firstName)
		fmt.Println("Last name:", emp1.lastName)
		fmt.Println("Age:", emp1.age)
		fmt.Printf("Salary: $%d\n", emp1.salary)
		emp1.salary = 700
		fmt.Printf("New alary: $%d\n", emp1.salary)
		fmt.Println()

		fmt.Println("Zero value of a struct")
		var emp4 Employees
		fmt.Println("First Name:", emp4.firstName)
		fmt.Println("Last Name:", emp4.lastName)
		fmt.Println("Age:", emp4.age)
		fmt.Println("Salary:", emp4.salary)
		fmt.Println()

		emp5 := Employees{
			firstName: "John",
			lastName: "Paul",
		}
		fmt.Println("First Name:", emp5.firstName)
		fmt.Println("Last Name:", emp5.lastName)
		fmt.Println("Age:", emp5.age)
		fmt.Println("Salary:", emp5.salary)
		fmt.Println()

		fmt.Println("Pointers to a struct")
		emp8 := &Employees{
			firstName: "Sam",
			lastName: "Anderson",
			age: 55,
			salary: 6000,
		}
		fmt.Println("First name:", (*emp8).firstName)
		fmt.Println("Age:", (*emp8).age)
		//This works too
		fmt.Println("First name:", emp8.firstName)
		fmt.Println("Age:", emp8.age)
		fmt.Println()

		fmt.Println("Anonymous fields")
		p1 := PersonAnon {
			string: "naveen",
			int: 50,
		}
		fmt.Println(p1.string)
		fmt.Println(p1.int)
		fmt.Println()

		fmt.Println("Nested structs")
		p2 := Person {
			name: "Naveen",
			age: 50,
			address : Address{
				city: "Chicago",
				state: "Illinois",
			},
		}
		fmt.Println("Name:", p2.name)
		fmt.Println("Age:", p2.age)
		fmt.Println("City:", p2.address.city)
		fmt.Println("State:", p2.address.state)
		fmt.Println()

		fmt.Println("Promoted fields")
		p3 := Person2 {
			name: "Naveen",
			age: 50,
			Address : Address{
				city: "Chicago",
				state: "Illinois",
			},
		}
		fmt.Println("Name:", p3.name)
		fmt.Println("Age:", p3.age)
		fmt.Println("City:", p3.city)
		fmt.Println("State:", p3.state)
		fmt.Println()

		fmt.Println("Exported structs and fields")
		spec := computer.Spec{
			Maker: "apple",
			Price: 50000,
		}
		fmt.Println("Maker:", spec.Maker)
		fmt.Println("Price:", spec.Price)
		fmt.Println()

		name1 := name{"Steve","Jobs"}
		name2 := name{"Steve","Johnson"}

		if name1 == name2 {
			fmt.Println("Name 1 and Name 2 are equal")
		} else {
			fmt.Println("Name 1 and Name 2 are not equal")
		}

		// IMPORTANT
		// Struct variables are not comparable if they contain fields that are not comparable(e.g. Maps)
	} else {
		// Do nothing
	}
}
