package subjects

import (
	"fmt"
)

type VowelsFinder interface {
	FindVowels() []rune
}

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empId int
	basicpay int
	pf int
}

type Contract struct {
	empId int
	basicpay int
}

type Freelancer struct {
	empId int
	ratePerHour int
	totalHours int
}

type Worker interface {
	Work()
}

type Person3 struct {
	name string
}

type Mystring string

func (p Permanent) CalculateSalary() int  {
	return p.basicpay + p.pf
}

func (c Contract) CalculateSalary() int  {
	return c.basicpay
}

func (f Freelancer) CalculateSalary() int {
	return f.ratePerHour * f.totalHours
}

func TotalExpense(s []SalaryCalculator)  {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}

func (p Person3) Work() {
	fmt.Println(p.name, "is working")
}

func describe(w Worker)  {
	fmt.Printf("Interface type %T value %v\n", w, w)
}

type Describer interface {
	Describe()
}

type Person4 struct {
	name string
	age int
}

func (p Person4) Describe() {
	fmt.Printf("%s is %d years old.", p.name, p.age)
}

func findType2(i interface{})  {
	switch v := i.(type){
	case Describer:
		v.Describe()
	default:
		fmt.Println("unknown type")
	}
}

func (ms Mystring) FindVowels() []rune  {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels,rune)
		}
	}
	return vowels
}

func describe2(i interface{})  {
	fmt.Printf("Type = %T, value %v\n", i, i)
}

func assert(i interface{})  {
	v,ok  := i.(int)
	fmt.Println(v, ok)
}

func findType(i interface{})  {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	default:
		fmt.Println("Unknown type")
	}
}

func PrintInterfacesPt1(print string)  {
	if print == "yes" {
		fmt.Println("Declaring and implementing an interface")
		name := Mystring("Sam Anderson")
		var v VowelsFinder
		v = name
		fmt.Printf("Vowels are %c", v.FindVowels())
		fmt.Println()
		fmt.Println()

		fmt.Println("Practical use of an interface")
		pemp1 := Permanent{1,5000,20}
		pemp2 := Permanent{2,6000,30}
		cemp1 := Contract{3,3000}
		freela1 := Freelancer{4, 70, 120}
		freela2 := Freelancer{5,100,100}
		employees := []SalaryCalculator{pemp1, pemp2, cemp1, freela1, freela2}
		TotalExpense(employees)
		fmt.Println()
		fmt.Println()

		fmt.Println("Interface internal representation")
		p := Person3{"Naveen"}
		var w Worker = p
		describe(w)
		w.Work()
		p.Work()
		fmt.Println()

		fmt.Println("Empty interface")
		s := "Hello World"
		describe2(s)
		i := 55
		describe2(i)
		strt := struct {
			name string
		} {
			name: "Raveen R",
		}
		describe2(strt)
		fmt.Println()

		fmt.Println("Type assertion")
		var s2 interface{} = 56
		assert(s2)
		fmt.Println(s2)

		var s3 interface{} = "Steven Two"
		assert(s3)
		fmt.Println()

		fmt.Println("Type switch")
		findType("Naveen")
		findType(77)
		findType(89.98)
		fmt.Println()

		// It is also possible to compare a type to an interface.
		// If we have a type and if that type implements an interface, it is possible to compare this type with the interface it implements.
		findType2("Naveen")
		p2 := Person4{"Naveen R", 25}
		findType2(p2)
	} else {
		// Do nothing
	}
}
