package subjects

import (
	"fmt"
)

type Describer2 interface {
	Describe()
}

type Person5 struct {
	name string
	age int
}

type Address2 struct {
	state string
	country string
}

type SalaryCalculator2 interface {
	DisplaySalary()
}

type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

type Employee3 struct {
	firstName string
	lastName string
	basicPay int
	pf int
	totalLeaves int
	leavesTaken int
}

type EmployeeOperations interface {
	SalaryCalculator2
	LeaveCalculator
}

func (e Employee3) DisplaySalary()  {
	fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee3) CalculateLeavesLeft() int  {
	return e.totalLeaves - e.leavesTaken
}

func (p Person5) Describe() {
	fmt.Println(p.name, "is", p.age, "years old.")
}

func (a *Address2) Describe()  {
	fmt.Println("State:", a.state, "Country:", a.country)
}



func PrintInterfacesPt2(print string)  {
	if print == "yes" {
		var d1 Describer
		p1 := Person5{"Sam",25}
		d1 = p1
		d1.Describe()

		p2 := Person5{"James",32}
		d1 = &p2
		d1.Describe()

		var d2 Describer
		a := Address2{"Washington","USA"}
		/* compilation error if the following line is
		   uncommented
		   cannot use a (type Address) as type Describer
		   in assignment: Address does not implement
		   Describer (Describe method has pointer
		   receiver)
		*/
		//d2 = a
		d2 = &a
		d2.Describe()
		fmt.Println()

		fmt.Println("Implementing multiple interfaces")
		e := Employee3{"Naveen", "Ramanathan", 5000, 200, 30, 5}
		var s SalaryCalculator2 = e
		s.DisplaySalary()
		var l LeaveCalculator = e
		l.CalculateLeavesLeft()
		fmt.Println()

		fmt.Println("Embedding interfaces")
		var empOp EmployeeOperations = e
		empOp.DisplaySalary()
		fmt.Println("\n Leave left =", empOp.CalculateLeavesLeft())
		fmt.Println()

		fmt.Println("Zero value of Interface")
		var d3 Describer
		if d3 == nil {
			fmt.Printf("dnil is nil and has type %T and value %v\n", d3, d3)
		}
		fmt.Println()
	} else {
		// Do nothing
	}
}
