package subjects

import "fmt"

func calculateBill(price, no int) int {
	// or calculate(price int, no int)
	var totalPrice = price * no
	return totalPrice
}

func rectProps2(length, width float64)(float64, float64) {  
    var area = length * width
    var perimeter = (length + width) * 2
    return area, perimeter
}

func rectProps(lenght, width float64)(area, perimeter float64) {
	area = lenght * width
	perimeter = (lenght + width) * 2
	return
	//area and perimeter are the named return values in the above function. 
	//Note that the return statement in the function does not explicitly return any value. 
	//Since area and perimeter are specified in the function declaration as return values, they are automatically returned from the function when a return statement in encountered.
}

func PrintFunctions(print string)  {
	if print == "yes" {
		price, no := 90, 6
		totalPrice := calculateBill(price, no)
		fmt.Println("Total price is", totalPrice,"\n")

		area, perimeter := rectProps(10.8, 5.6)
		fmt.Printf("Area %f Perimeter %f", area, perimeter)
		fmt.Printf("\n")

		area2, _ := rectProps(10.8, 5.6) // perimeter is discarted
		fmt.Printf("Area %f", area2)
	} else {
		// Do nothing
	}
}