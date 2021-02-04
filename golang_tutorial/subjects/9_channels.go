package subjects

import (
	"fmt"
	"time"
)

func hello3(done chan bool){
	fmt.Println("Hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("Hello go routine awake and going to write to done")
	done <- true
}

func digits(number int, dchn1 chan int) {
	for number != 0 {
		digit := number % 10
		dchn1 <- digit
		number /= 10
	}
	close(dchn1)
}

func calcSquares(number int, squareop chan int)  {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int)  {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeop <- sum
}

func sendData(sendch chan<- int) {
	sendch <- 10
}

func producer(chn1 chan int) {
	for i := 0; i < 10; i++ {
		chn1 <- i

	}
	close(chn1)
}

func PrintChannels(print string)  {
	if print == "yes" {
		fmt.Println("Declaring channels")
		var a chan int
		if a == nil {
			fmt.Println("Channel a is nil, going to define it")
			a = make(chan int)
			fmt.Printf("Type of a is %T\n", a)
		}
		fmt.Println()

		fmt.Println("Channel example program")
		done := make(chan bool)
		fmt.Println("Main going to call hello go routine")
		go hello3(done)
		<- done
		fmt.Println("main received function")
		fmt.Println()

		fmt.Println("Another example for channels")
		number := 589
		sqrch := make(chan int)
		cubech := make(chan int)
		go calcCubes(number, sqrch)
		go calcSquares(number, cubech)
		squares, cubes := <- sqrch, <- cubech
		fmt.Println("Final output:", squares + cubes)
		fmt.Println()

		chn1 := make(chan int)
		go sendData(chn1)
		fmt.Println(<-chn1)
		fmt.Println()

		fmt.Println("Closing channels and for range loops on channels")
		ch := make(chan int)
		go producer(ch)
		for v := range ch {
			fmt.Println("Received", v)
		}
	} else {}
	// Do nothing
}
