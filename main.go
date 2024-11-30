package main

import (
	"goroutines-patterns/generator"
	producerconsumer "goroutines-patterns/producer-consumer"
)

func main() {
	generatorPattern()

}

func generatorPattern() {
	execute(generator.ExamplePattern{})

	execute(generator.LogGenerator{})

	execute(generator.EcommerceExample{})

	execute(generator.DataPointGenerator{})

	execute(generator.StreamGenerator{})

	execute(generator.LargeSetPattern{})
}

func producerConsumerPattern() {
	execute(producerconsumer.ProducerConsumerPattern{})

	execute(producerconsumer.RealWorldPattern{})

	execute(producerconsumer.ProducerConsumerAdvanced{})
}

func execute(patterns Patterns) {
	patterns.Execute()
}
