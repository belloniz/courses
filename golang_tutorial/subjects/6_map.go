package subjects

import (
	"fmt"
)

type employee struct {
	salary int
	country string
}

func PrintMap(print string)  {
	if print == "yes" {
		employeeSalary := make(map[string]int)
		fmt.Println(employeeSalary)
		employeeSalary["steve"] = 1200
		employeeSalary["jamie"] = 1500
		employeeSalary["mike"] = 9000
		fmt.Println("EmployeeSalary map contents:", employeeSalary)
		fmt.Println()

		// Other option
		employeeSalary2 := map[string]int {
			"steve": 1500,
			"jamie": 1200,
		}
		employeeSalary2["mike"] = 900
		fmt.Println("EmployeeSalary map contents:", employeeSalary2)
		fmt.Println()

		employee1 := "jamie"
		salary := employeeSalary[employee1]
		fmt.Println("Salary of", employee1, "is", salary)
		fmt.Println()

		newEmp := "oliver"
		value, ok := employeeSalary[newEmp]
		if ok == true {
			fmt.Println("Salary of", newEmp, "is", value)
			return
		}
		fmt.Println(newEmp, "not found")
		fmt.Println()

		employeeSalary3 := map[string]int {
			"steve": 12000,
			"jamie": 15000,
			"mike": 9000,
		}
		fmt.Println("Contents of the map")
		for key, value := range employeeSalary3{
			fmt.Println("Employee Salary", key, "is", value)
		}
		fmt.Println()

		fmt.Println("map before deletion", employeeSalary3)
		delete(employeeSalary3, "steve")
		fmt.Println("map after deletion", employeeSalary3)
		fmt.Println()

		emp1 := employee{
			salary:  12000,
			country: "USA",
		}

		emp2 := employee {
			salary: 14000,
			country: "Canada",
		}

		emp3 := employee {
			salary : 13000,
			country: "India",
		}

		employeeInfo := map[string]employee{
			"Steve": emp1,
			"Jamie": emp2,
			"Mike": emp3,
		}

		for name, info := range employeeInfo {
			fmt.Printf("Employee: %s Salary: $%d Country:%s\n",name, info.salary, info.country)
		}

		fmt.Println("legth of EmployeeInfo is",len(employeeInfo))
		fmt.Println()

		employeeSalary4 := map[string]int{
			"steve": 12000,
			"jamie": 15000,
			"mike": 9000,
		}

		//employeeSalary5 := map[string]int{
		//	"steve": 12000,
		//	"jamie": 15000,
		//	"mike": 8000,
		//}

		fmt.Println("Original employee salary", employeeSalary4)
		modified := employeeSalary4
		modified["mike"] = 18000
		fmt.Println("Employee salary changed",employeeSalary4)
		fmt.Println()

		for i, v := range employeeSalary4 {
			fmt.Println(i)
			fmt.Println(v)
		}
	} else {
		// Do nothing
	}
}
