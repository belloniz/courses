package subjects

import (
	"fmt"
	"time"
	"sync"
)

func write(ch chan int)  {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("Sucessfully wrote", i, "to ch.")
	}
	close(ch)
}

func process(i int, wg * sync.WaitGroup)  {
	fmt.Println("Started Goroutine", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

func PrintBufferedChannelsAndWorkerPools(print string)  {
	if print == "yes" {
		fmt.Println("Example")
		ch := make(chan string, 2)
		ch <- "naveen"
		ch <- "paul"
		fmt.Println(<- ch)
		fmt.Println(<- ch)
		fmt.Println()

		fmt.Println("Another example")
		ch2 := make(chan int, 2)
		go write(ch2)
		time.Sleep(2 * time.Second)
		for v := range ch2 {
			fmt.Println("read value", v, "from ch")
			time.Sleep( 2 * time.Second)
		}
		fmt.Println()

		fmt.Println("Deadlock")
		//ch3 := make(chan string, 2)
		//ch3 <- "naveen"
		//ch3 <- "paul"
		//ch3 <- "steve"
		//fmt.Println(<- ch)
		//fmt.Println(<- ch)
		fmt.Println()

		fmt.Println("Length vs Capacity")
		ch4 := make(chan string, 3)
		ch4 <- "naveen"
		ch4 <- "paul"
		fmt.Println("capacity is", cap(ch4))
		fmt.Println("length is", len(ch4))
		fmt.Println("read value", <- ch4)
		fmt.Println("new length is", len(ch4))
		fmt.Println()

		fmt.Println("WaitGroup")
		no := 3
		var wg sync.WaitGroup
		for i := 0; i < no; i++ {
			wg.Add(1)
			go process(i, &wg)
		}
		wg.Wait()
		fmt.Println("All go routines finished executing.")
	} else {
		// Do nothing
	}
}
