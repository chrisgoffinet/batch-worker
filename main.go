package main

import (
	"fmt"
	"sync"
)

// process handles the actual computation
func process(batch []int) {
	fmt.Println(batch)
}

// worker handles batching up events until batchSize is reached then calls process()
func worker(wg *sync.WaitGroup, bufCh chan int, batchSize int) {
	var batch []int
	defer wg.Done()

	for {
		select {
		case v, open := <-bufCh:
			if open == false {
				// anything left to process, finish it
				if len(batch) > 0 {
					process(batch)
				}
				return
			}
			batch = append(batch, v)
			if len(batch) >= batchSize {
				process(batch)
				batch = batch[:0]
			}
		}
	}
}

func main() {
	// setup wait group so we can wait for all goroutines to complete
	wg := new(sync.WaitGroup)
	batchSize := 10

	// create a bounded channel that will block to minimize memory
	bufCh := make(chan int, batchSize)

	// spawn enough workers as goroutines
	for index := 0; index < 5; index++ {
		fmt.Printf("started worker: %d\n", index)
		go worker(wg, bufCh, batchSize)
		wg.Add(1)
	}

	// add some integers to the channel to have batched and processed
	for index := 0; index < 999; index++ {
		bufCh <- index
	}
	close(bufCh)
	wg.Wait()
}
