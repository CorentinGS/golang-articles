package generator

import (
	"fmt"
	"time"
)

type StreamGenerator struct{}

type DataItem struct {
	ID   int
	Data string
}

// mockDataStream simulates a data source (e.g., a file, queue, or network stream)
func mockDataStream(count int) <-chan DataItem {
	out := make(chan DataItem)
	go func() {
		defer close(out)
		for i := 0; i < count; i++ {
			// Simulate reading from a data source
			time.Sleep(100 * time.Millisecond)
			out <- DataItem{
				ID:   i + 1,
				Data: fmt.Sprintf("Data-%d", i+1),
			}
		}
	}()
	return out
}

// dataGenerator consumes the mock stream and yields processed data
func dataGenerator(stream <-chan DataItem) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for item := range stream {
			// Process the data item
			processedData := fmt.Sprintf("Processed: %s (ID: %d)", item.Data, item.ID)
			out <- processedData
		}
	}()
	return out
}

func (g StreamGenerator) Execute() {
	// Create a mock data stream
	dataStream := mockDataStream(10)

	// Create a generator to process the stream
	processedDataGen := dataGenerator(dataStream)

	// Consume and print the processed data
	for data := range processedDataGen {
		fmt.Println(data)
	}
}
