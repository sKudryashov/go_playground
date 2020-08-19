package flowpanic

import (
	"fmt"
)

func InitFlow() {
	fmt.Println("Init flow defer panic ...")
	flowDefer()
	fmt.Println("Init flow finished, recovered here")
	// cfg := api.Config{}
	// // cl, err := &api.NewClient(cfg)
	// if err != nil {
	// 	println("err ", err.Error())
	// }
	// println("trpr", cl)
}

func flowDefer() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in flow defer with r: ", r)
		}
	}()
	fmt.Println("Calling flowPanic start.. ")
	flowPanic(0) //#1
	// if r := recover(); r != nil {
	// 	fmt.Println("Recovered in flow defer with r: ", r)
	// }
	flowPanicEmbedDefer(0) //#2
	fmt.Println("Returned normally from flowPanic")
}

func flowPanic(i int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("flowPanic r: ", r)
		}
	}()
	if i > 1 {
		fmt.Println("Let's make a small panic... Panicking!!!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in flow panic, i=", i)
	fmt.Println("Usual flow in flowPanic, i=", i)
	flowPanic(i + 1)
}

func flowPanicEmbedDefer(i int) {
	if i > 3 {
		fmt.Println("Let's make a small panic... Panicking!!!")
		panic(fmt.Sprintf("%v", i))
	}
	defer func() {
		if re := recover(); re != nil {
			fmt.Println("Recovering with flowPanicEmbedDefer", re)
		}
	}()
	fmt.Println("Usual flow in flowPanic, i=", i)
	flowPanicEmbedDefer(i + 1)
	fmt.Println("Recovered embedded defer?") // it is hit in case #2
}
