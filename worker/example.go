package worker

import (
	"fmt"
	"time"
)

type PoolPattern struct{}

func (w PoolPattern) Execute() {
	const numJobs = 5
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start worker pool
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs to the workers
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect and print results
	for a := 1; a <= numJobs; a++ {
		result := <-results
		fmt.Printf("Job result: %d\n", result)
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Second) // Simulating work
		results <- job * 2
	}
}
