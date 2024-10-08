package main

import (
	"log"
	"math/rand"
	"sync/atomic"
)

const totalSimulations uint64 = 1_000_000

var completedSimulations uint64

func roll() int {
	var count int
	for i := 1; i <= 231; i++ {
		if rand.Intn(4) == 0 {
			count++
		}
	}
	atomic.AddUint64(&completedSimulations, 1)
	return count
}

func worker(jobs <-chan int, results chan<- int) {
	for range jobs {
		results <- roll()
	}
}

func main() {
	const pool = 12
	const upper = int(totalSimulations)

	jobs := make(chan int, upper)
	results := make(chan int, upper)

	for w := 1; w <= pool; w++ {
		go worker(jobs, results)
	}

	for j := 1; j <= upper; j++ {
		jobs <- j
	}
	close(jobs)

	var max int
	for n := 1; n <= upper; n++ {
		output := <-results
		if output > max {
			max = output
		}
	}

	if completedSimulations != totalSimulations {
		log.Fatalf("Expected %d simulations but instead only %d were performed", totalSimulations, completedSimulations)
	} else {
		println(max)
	}
}
