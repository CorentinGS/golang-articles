package generator

import (
	"fmt"
	"time"
)

type SetDataItem struct {
	Data string
	ID   int
}

// lazyDataLoader simulates loading a large dataset lazily
func lazyDataLoader(filePath string) <-chan SetDataItem {
	out := make(chan SetDataItem)
	go func() {
		defer close(out)
		// Simulate opening a large file
		fmt.Printf("Opening file: %s\n", filePath)

		// Simulate reading the file line by line
		for i := 0; i < 1000000; i++ {
			// Simulate processing delay for each item
			time.Sleep(1 * time.Millisecond)
			out <- SetDataItem{
				ID:   i + 1,
				Data: fmt.Sprintf("Data from line %d", i+1),
			}
			if i%100000 == 0 {
				fmt.Printf("Processed %d items\n", i)
			}
		}
	}()
	return out
}

func processData(data <-chan SetDataItem) {
	for item := range data {
		// Simulate data processing
		processedData := fmt.Sprintf("Processed: %s (ID: %d)", item.Data, item.ID)
		fmt.Println(processedData)
	}
}

type LargeSetPattern struct{}

func (set LargeSetPattern) Execute() {
	dataStream := lazyDataLoader("large_dataset.txt")
	processData(dataStream)
}
