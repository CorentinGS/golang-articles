package worker

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

type ImageProcessor struct{}

func (i ImageProcessor) Execute() {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	numWorkers := numCPU * 2 // Use 2 workers per CPU core
	const numJobs = 10

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	// Initialize worker pool
	for w := 1; w <= numWorkers; w++ {
		go imageProcessor(w, jobs, results)
	}

	// Send image processing jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{
			ID:       j,
			ImageURL: fmt.Sprintf("https://example.com/image%d.jpg", j),
			Size:     100 * j, // Varying sizes
		}
	}
	close(jobs)

	// Collect and handle results
	for a := 1; a <= numJobs; a++ {
		result := <-results
		if result.Error != nil {
			fmt.Printf("Error processing image %d: %v\n", result.JobID, result.Error)
		} else {
			fmt.Printf("Successfully processed image %d to size %dpx in %v\n",
				result.JobID, result.NewSize, result.TimeSpent)
		}
	}
}

type Job struct {
	ImageURL string
	ID       int
	Size     int
}

type Result struct {
	Error     error
	ImageURL  string
	JobID     int
	NewSize   int
	TimeSpent time.Duration
}

func imageProcessor(id int, jobs <-chan Job, results chan<- Result) {
	for job := range jobs {
		startTime := time.Now()

		fmt.Printf("Worker %d processing image %d from %s\n", id, job.ID, job.ImageURL)

		result := Result{
			JobID:    job.ID,
			ImageURL: job.ImageURL,
			NewSize:  job.Size,
		}

		// Simulate image processing with realistic steps
		err := processImage(job)
		if err != nil {
			result.Error = err
			results <- result
			continue
		}

		result.TimeSpent = time.Since(startTime)
		results <- result
	}
}

func processImage(job Job) error {
	// Simulate various image processing steps
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

	// Simulate potential errors
	if rand.Float32() < 0.1 {
		return fmt.Errorf("failed to process image %d: simulation error", job.ID)
	}

	return nil
}
