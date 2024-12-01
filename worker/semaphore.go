package worker

import (
	"fmt"
	"sync"
	"time"
)

type SemaphoreWorker struct{}

func (s SemaphoreWorker) Execute() {
	tasks := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	processWithSemaphore(tasks, 3)
}

func processWithSemaphore(tasks []int, maxConcurrency int) {
	sem := make(chan struct{}, maxConcurrency)
	var wg sync.WaitGroup

	for _, task := range tasks {
		wg.Add(1)
		sem <- struct{}{} // Acquire semaphore
		go func(task int) {
			defer wg.Done()
			defer func() { <-sem }() // Release semaphore
			processTask(task)
		}(task)
	}

	wg.Wait()
}

func processTask(task int) {
	fmt.Printf("Processing task %d\n", task)
	time.Sleep(time.Second) // Simulate work
	fmt.Printf("Completed task %d\n", task)
}
