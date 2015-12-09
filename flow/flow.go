package flow

import "fmt"

func InitFlow()  {
	fmt.Println("Init flow defer panic ...");
	flowDefer()
	fmt.Println("Init flow finished")
}

func flowDefer()  {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in flow defer with r: ", r)
		}
	} ()
	fmt.Println("Calling flowPanic start.. ");
	flowPanic(0) //#1
	flowPanicEmbedDefer(0)//#2
	fmt.Println("Returned normally from flowPanic")
}

func flowPanic(i int)  {
	if i > 3 {
		fmt.Println("Let's make a small panic... Panicking!!!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in flow panic, i=", i)
	fmt.Println("Usual flow in flowPanic, i=", i)
	flowPanic(i + 1)
}

func flowPanicEmbedDefer(i int)  {
	if i > 3 {
		fmt.Println("Let's make a small panic... Panicking!!!")
		panic(fmt.Sprintf("%v", i))
	}
	defer func() {
		if re := recover(); re!=nil{
			fmt.Println("Recovering with flowPanicEmbedDefer", re)
		}
	} ()
	fmt.Println("Usual flow in flowPanic, i=", i)
	flowPanicEmbedDefer(i + 1)
	fmt.Println("Recovered embedded defer?") // it is hit in case #2
}
