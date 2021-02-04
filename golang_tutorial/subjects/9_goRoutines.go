package subjects

import (
	"fmt"
	"time"
)

func hello2(){
	fmt.Println("Hello World Routine")
}

func numbers() {
	for i:= 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func alphabets()  {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}

}

func PrintGoRoutines(print string)  {
	if print == "yes" {
		fmt.Println("How to start a Goroutine?")
		go hello2()
		time.Sleep(1 * time.Second)
		fmt.Println("main function")
		fmt.Println()

		fmt.Println("Starting multiple Goroutines")
		go numbers()
		go alphabets()
		time.Sleep(3000 * time.Millisecond)
		fmt.Println("main terminated")
	} else {
		// Do nothing
	}
}