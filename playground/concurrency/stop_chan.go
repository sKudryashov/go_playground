package concurrency

import (
	"fmt"
	"time"
)

func CloseStopTest() {
	ch := make(chan interface{})
	go worker1(ch)
	go worker2(ch)
	go worker3(ch)

	go worker3(ch)
	go worker3(ch)
	go worker3(ch)

	go worker3(ch)
	go worker3(ch)
	go worker3(ch)

	go worker3(ch)
	// if each worker performs nil chan than it seems to be guaranteed that each one gets signal which works only once
	// 	worker2 stopped
	// worker3 stopped
	// worker1 stopped
	// tested on 10 workers
	close(ch) // all three workers read the channel but multiple times
	time.Sleep(5 * time.Second)
}

func LaunchStopTest() {
	ch := make(chan interface{})
	ch2 := make(chan interface{})
	ch4 := make(chan interface{})
	go worker1(ch)
	go worker2(ch)
	go worker3(ch2)
	go worker4(ch4)

	close(ch2)
	time.Sleep(3 * time.Second)
	ch <- struct{}{} // stops one worker
	ch <- struct{}{} // stops two workers
	close(ch)        // after closing channel and not nil-ing it i constanly get this messages in select lile "worker2 stopped"
	// so after channel closed we have to nil it
	fmt.Println("finish")
	time.Sleep(1 * time.Second)
	ch4 = nil
	ch4 <- struct{}{} // expect panic here - no panic but blocked forever
}

func worker1(chStop chan interface{}) {
	for {
		select {
		case <-chStop:
			chStop = nil
			fmt.Println("worker1 stopped")
		}
	}
}

func worker4(chStop chan interface{}) {
	for {
		select {
		case <-chStop:
			fmt.Println("worker4 doesnt have to be stopped, it's channel nilled")
			// when close the chan -  i get the signal chan closed
		}
	}
}

func worker3(chStop chan interface{}) {
	for {
		select {
		case <-chStop:
			chStop = nil
			fmt.Println("worker3 stopped")
			// when close the chan -  i get the signal chan closed
			chStop = nil // otherwise constantly get "worker3 stopped"
		}
	}
}

func worker2(chStop chan interface{}) {
	for {
		select {
		case <-chStop:
			chStop = nil
			fmt.Println("worker2 stopped")
		}
	}
}
