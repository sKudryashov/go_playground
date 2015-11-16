package concurrency

import "fmt"

func Launch () {
	SetBufferedWriter()
}

func SetUnbufferedWriter () {
	unbuffCh := make(chan int)
	i := 0
	iout := 0
	var totalSum int

	go func () {
		for {
			i++
			if (i > 1000) {
				break
			}
		}
		unbuffCh <- i
	} ()

	for {
		iout++
		if (iout > 1000) {
			break
		}
	}
	totalSum <- unbuffCh

	fmt.Println("Total concurrency sum:", totalSum)
	fmt.Println("Total iout sum:", iout)
}

func SetBufferedWriter () {

}