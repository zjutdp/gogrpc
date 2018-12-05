package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

const jobCount = 10000000

var totalConsumed uint64

func main() {

	start := time.Now()

	// Similar to bufio, buffered channel speed up
	// communication between goroutings
	jobs := make(chan int, 1000)
	results := make(chan int, 1000)
	done := make(chan bool)

	workerCount := 30
	consumerCount := 50

	for i := 0; i < workerCount; i++ {
		go worker(i, jobs, results)
	}
	fmt.Println("Started ", workerCount, " workers")

	for i := 0; i < consumerCount; i++ {
		go consumer(i, results, done)
	}
	fmt.Println("Started ", consumerCount, " consumer")

	for i := 0; i < jobCount; i++ {
		jobs <- i
	}
	fmt.Printf("Generated %d jobs\n", jobCount)

	<-done
	fmt.Println("All jobs done and all results consumed in ",
		time.Now().Sub(start))
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		//fmt.Println("Worker: ", id, " processing job: ", j)

		//time.Sleep(100 * time.Millisecond)

		results <- j * 10
	}
}

func consumer(id int, results <-chan int, done chan<- bool) {
	for r := range results {
		//fmt.Println("Consumer: ", id, " consuming result: ", r)
		r++
		//time.Sleep(100 * time.Millisecond)

		atomic.AddUint64(&totalConsumed, 1)
		if atomic.LoadUint64(&totalConsumed) == jobCount {
			done <- true
		}
	}
	// Same functionality as above code fragment but more complex
	// for {
	// 	r, ok := <-results
	// 	if !ok {
	// 		return
	// 	}
	// 	fmt.Println("Consumer: ", id, " consuming result: ", r)
	// }
}
