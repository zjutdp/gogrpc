package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const jobCount = 10000000

var totalConsumed uint64
var totalProduced uint64

func main() {

	start := time.Now()

	// Similar to bufio, buffered channel speed up
	// communication between goroutings by buffering
	jobs := make(chan int, 1000)
	results := make(chan int, 1000)

	wgWorker := new(sync.WaitGroup)
	wgConsumer := new(sync.WaitGroup)

	done := make(chan bool, 1) // Has to be 1, otherwise deadlock

	workerCount := 30
	consumerCount := 50

	for i := 0; i < workerCount; i++ {
		wgWorker.Add(1)
		go worker(i, jobs, results, wgWorker)
	}
	fmt.Println("Started ", workerCount, " workers")

	for i := 0; i < consumerCount; i++ {
		wgConsumer.Add(1)
		go consumer(i, results, done, wgConsumer)
	}
	fmt.Println("Started ", consumerCount, " consumers")

	for i := 0; i < jobCount; i++ {
		jobs <- i
	}
	fmt.Printf("Generated %d jobs\n", jobCount)

	close(jobs)
	fmt.Println("Close jobs channel as jobs done")

	wgWorker.Wait()
	fmt.Printf("All workers have finished with total: %d produced!\n",
		atomic.LoadUint64(&totalProduced))

	close(results)
	fmt.Println("Close results channel as workers done, no more write")

	wgConsumer.Wait()
	fmt.Printf("All consumers have finished with total: %d consumed!\n",
		atomic.LoadUint64(&totalConsumed))

	<-done // Same effect as above all consumers have finished waiting

	// Only OK when "done" is closed by the sender, otherwise deadlock
	// Adding below for closing this channel
	if _, ok := <-done; !ok {
		fmt.Println("All jobs done and all results consumed in ",
			time.Now().Sub(start))
	}

	// Double check
	if len(done) == 0 {
		fmt.Println("Double checked that all jobs done and all results consumed!")
	}
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	for j := range jobs {
		//fmt.Println("Worker: ", id, " processing job: ", j)
		//time.Sleep(100 * time.Millisecond)
		atomic.AddUint64(&totalProduced, 1)
		results <- j * 10
	}

	wg.Done()
}

func consumer(id int, results <-chan int, done chan<- bool, wg *sync.WaitGroup) {
	for r := range results {
		//fmt.Println("Consumer: ", id, " consuming result: ", r)
		r++
		//time.Sleep(100 * time.Millisecond)
		atomic.AddUint64(&totalConsumed, 1)
		if atomic.LoadUint64(&totalConsumed) == jobCount {
			done <- true
			close(done)
		}
	}

	wg.Done()
	// Same functionality as above "range" code fragment but more complex
	// for {
	// 	r, ok := <-results
	// 	if !ok {
	// 		return
	// 	}
	// 	fmt.Println("Consumer: ", id, " consuming result: ", r)
	// }
}
