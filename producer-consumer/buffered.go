package producer_consumer

import (
	"fmt"
	"math/rand"
	"time"
)

type ProducerConsumerAdvanced struct{}

type LogEntry struct {
	Timestamp time.Time
	Level     string
	Message   string
}

func logGenerator(logs chan<- LogEntry) {
	// Simulate incoming logs
	for i := 0; i < 10; i++ {
		log := LogEntry{
			Timestamp: time.Now(),
			Level:     []string{"INFO", "WARNING", "ERROR"}[rand.Intn(3)],
			Message:   fmt.Sprintf("Event #%d occurred", i),
		}
		logs <- log
		time.Sleep(100 * time.Millisecond) // Simulate varying log frequencies
	}
	close(logs)
}

func logProcessor(logs <-chan LogEntry) {
	batch := make([]LogEntry, 0, 3) // Process logs in batches of 3

	for log := range logs {
		batch = append(batch, log)

		if len(batch) == 3 {
			// Process batch
			processLogBatch(batch)
			batch = batch[:0] // Clear the batch
		}
	}

	// Process remaining logs
	if len(batch) > 0 {
		processLogBatch(batch)
	}
}

func processLogBatch(batch []LogEntry) {
	fmt.Println("Processing batch of logs:")
	for _, log := range batch {
		fmt.Printf("[%s] %s: %s\n",
			log.Timestamp.Format("15:04:05"),
			log.Level,
			log.Message)
	}
	fmt.Println("Batch processing complete\n")
}

func (p ProducerConsumerAdvanced) Execute() {
	logs := make(chan LogEntry)
	go logGenerator(logs)
	logProcessor(logs)
}
