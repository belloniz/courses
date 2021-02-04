package main

import (
	"golang_tutorial/subjects"
)

type employee struct {
	salary  int
	country string
}

func main()  {
	// Part 1
	subjects.PrintIntroduction("no")

	// Part 2
	subjects.PrintConstants("no")
	subjects.PrintTypes("no")
	subjects.PrintVariables("no")

	// Part 3
	subjects.PrintFunctions("no")

	// Part 4
	subjects.PrintIfElseStatement("no")
	subjects.PrintLoops("no")
	subjects.PrintSwitchStatement("no")

	// Part 5
	subjects.PrintArrays("no")
	subjects.PrintSlices("no")
	subjects.PrintVariadicFunctions("no")

	//Part 6
	subjects.PrintMap("no")
	subjects.PrintStrings("no")

	//Part 7
	subjects.PrintPointers("no")
	subjects.PrintStructs("no")
	subjects.PrintMethods("no")

	//Part 8
	subjects.PrintInterfacesPt1("no")
	subjects.PrintInterfacesPt2("no")

	// Part 9
	subjects.PrintGoRoutines("no")
	subjects.PrintChannels("no")
	subjects.PrintBufferedChannelsAndWorkerPools("no")
	subjects.PrintBufferedChannelsAndWorkerPoolspt2("no")
}