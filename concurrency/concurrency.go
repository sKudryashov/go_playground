package concurrency

import (
	"fmt"
	"github.com/sKudryashov/go_playground/methods"
	"reflect"
)

func Launch () {
	SetUnbufferedWriter()
	SetBufferedWriter()
}

type GreetingsSet []methods.GreetingsManager

func SetUnbufferedWriter () {
	unbuffCh := make(chan int)
	i := 0
	iout := 0
	var totalSum bool

	go func () {
		for {
			i++
			if (i > 1000) {
				break
			}

		}
		unbuffCh <- i
		close(unbuffCh)
	} ()

	for {
		iout++
		if (iout > 1000) {
			break
		}
	}

	//#time.Sleep(time.Duration(100))
	 for c := range unbuffCh {
		 fmt.Println("Channel returned:", c)
	}

	fmt.Println("Total concurrency sum:", totalSum)
	fmt.Println("Total iout sum:", iout)
}

func SetBufferedWriter () {
	manager := GreetingsSet{
		{"Ukraine", "Vitaemo!", 2},
		{"English", "How is it going?", 0},
		{"USA", "How r u?", 8},
	}

	buffCh := make(chan methods.GreetingsManager, 30)

	//Both variants are working whether we use routine or just work with a channel
//	go func () {
//		for _, gr := range manager {
//			fmt.Println("gr.Text", reflect.TypeOf(gr.Text))
//			buffCh <- gr
//		}
//	}()
	for _, gr := range manager {
		fmt.Println("gr.Text", reflect.TypeOf(gr.Text))
		buffCh <- gr
	}
	close(buffCh)
	fmt.Println("buffCh", reflect.TypeOf(buffCh))

	for channelElement := range buffCh {
		fmt.Printf("Channel data: Country - CHannel Element: ", channelElement.Country, channelElement.Text)
	}

	managerAsia := GreetingsSet{
		{"Singapore", "How is it going?", 10},
		{"South Korea", "How is it going?", 12},
		{"Hong Kong", "How r u?", 11},
	}

    setTwoParallelWritersAndGetFirst(manager, managerAsia)
}

func setTwoParallelWritersAndGetFirst (setOne, setTwo GreetingsSet) {

	chan1 := make(chan string, 5)
	chan2 := make(chan string, 5)

	go func () {
		for _, item := range setOne {
			chan1 <- fmt.Sprintln("%s", item.Country)
		}
		close(chan1)
	}()
	go func () {
		for _, item := range setTwo {
			chan2 <- fmt.Sprintln("%s", item.Country)
		}
		close(chan2)
	}()

	for {
		select {
		case data, ok := <- chan2:
			if ok {
				fmt.Println("The state chan2 is:", ok, " and the data is: ", data)
			} else {
				return
			}
		case data, ok := <- chan1:
			if ok {
				fmt.Println("The state is chan1 :", ok, " and the data is: ", data)
			} else {
				return
			}
		default:
			fmt.Println("Waiting...")
		}
	}
}