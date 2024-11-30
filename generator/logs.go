package generator

import (
	"fmt"
	"math/rand"
	"time"
)

type LogGenerator struct{}

type LogEntry struct {
	Timestamp time.Time
	Level     string
	Message   string
}

func logGenerator(count int) <-chan LogEntry {
	out := make(chan LogEntry)
	go func() {
		levels := []string{"INFO", "WARNING", "ERROR"}
		messages := []string{
			"User logged in",
			"Failed login attempt",
			"Database connection lost",
			"API request received",
			"Cache miss",
		}
		for i := 0; i < count; i++ {
			out <- LogEntry{
				Timestamp: time.Now().Add(time.Duration(i) * time.Second),
				Level:     levels[rand.Intn(len(levels))],
				Message:   messages[rand.Intn(len(messages))],
			}
			time.Sleep(100 * time.Millisecond) // Simulate delay between log entries
		}
		close(out)
	}()
	return out
}

func (g LogGenerator) Execute() {
	for entry := range logGenerator(5) {
		fmt.Printf("[%s] %s: %s\n", entry.Timestamp.Format(time.RFC3339), entry.Level, entry.Message)
	}
}
