package concurrency

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)

// based on https://medium.com/@deckarep/gos-extended-concurrency-semaphores-part-1-5eeabfa351ce

func LaunchWeightedBlockingSemaphore() {
	wsem := semaphore.NewWeighted(int64(3))
	ctx := context.TODO()
	rr := []int{1, 2, 3, 4, 5, 5, 6, 7}
	for _, item := range rr {
		if err := wsem.Acquire(ctx, 1); err != nil {
			fmt.Println("error acquiring context", err.Error())
		}
		go func(employeeNum int) {
			fmt.Println("employee num ", employeeNum)
			time.Sleep(2 * time.Second)
			wsem.Release(1)
		}(item)
	}
}

func LaunchBlockingSemaphore() {
	sem := make(chan bool, 10)
	t := 0
	for i := 0; i < 20; i++ {
		sem <- true
		t++
		go Worker(t, sem)
	}
}

func Worker(index int, chDone chan bool) {
	fmt.Println("worker launched ", index)
	time.Sleep(5 * time.Second)
	<-chDone
	return
}
