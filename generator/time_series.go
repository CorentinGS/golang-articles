package generator

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type DataPoint struct {
	Timestamp time.Time
	Value     float64
}

func timeSeriesGenerator(
	duration time.Duration,
	interval time.Duration,
	baseValue float64,
	volatility float64,
) <-chan DataPoint {
	out := make(chan DataPoint)
	go func() {
		// Calculate number of points to generate
		startTime := time.Now()
		currentValue := baseValue

		// Generate points until duration is reached
		for t := startTime; t.Before(startTime.Add(duration)); t = t.Add(interval) {
			// Add random walk to the value
			change := (rand.Float64() - 0.5) * volatility
			currentValue += change

			out <- DataPoint{
				Timestamp: t,
				Value:     currentValue,
			}
		}
		close(out)
	}()
	return out
}

type DataPointGenerator struct{}

func (g DataPointGenerator) Execute() {
	for point := range timeSeriesGenerator(10*time.Second, time.Second, 100.0, 0.1) {
		fmt.Printf("[%s] %.2f\n", point.Timestamp.Format(time.Stamp), point.Value)
	}
}
