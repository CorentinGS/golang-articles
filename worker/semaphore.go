package worker

import (
	"fmt"
	"time"
)

type SemaphoreWorker struct{}

func (s SemaphoreWorker) Execute() {

	tasks := make([]Task, 10)
	for i := 0; i < len(tasks); i++ {
		tasks[i] = Task{
			ID:   i,
			Name: "Task " + string(i),
		}
	}

	results := processWithSemaphore(tasks, 3)

	for _, result := range results {
		println("Task ID:", result.TaskID, "Status:", result.Status)
	}

}

type Task struct {
	ID   int
	Name string
}

type TaskResult struct {
	TaskID int
	Status string
}

func processWithSemaphore(tasks []Task, maxConcurrency int) []TaskResult {
	sem := make(chan struct{}, maxConcurrency)
	results := make([]TaskResult, len(tasks))

	for i, task := range tasks {
		sem <- struct{}{} // Acquire semaphore
		go func(i int, task Task) {
			defer func() { <-sem }() // Release semaphore
			fmt.Printf("Processing task %d\n", task.ID)
			results[i] = processTask(task)
		}(i, task)
	}

	// Wait for all goroutines to finish
	for i := 0; i < maxConcurrency; i++ {
		sem <- struct{}{}
	}

	return results
}

func processTask(task Task) TaskResult {
	// Simulate processing time
	time.Sleep(1 * time.Second)
	return TaskResult{
		TaskID: task.ID,
		Status: "Completed",
	}
}
