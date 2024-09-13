package main

import "math/rand"

func roll() int {
	var count int
	for i := 1; i <= 231; i++ {
		if rand.Float32() <= 0.25 {
			count++
		}
	}
	return count
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for range jobs {
		results <- roll()
	}
}

func main() {
	const pool = 8
	const upper = 1_000_000_000
	rolls := upper / pool

	jobs := make(chan int, rolls)
	results := make(chan int, rolls)

	for w := 1; w <= pool; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= rolls; j++ {
		jobs <- j
	}
	close(jobs)

	var max int
	for n := 1; n <= rolls; n++ {
		output := <-results
		if output > max {
			max = output
		}
	}

	print(max)
}
