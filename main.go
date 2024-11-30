package main

import producerconsumer "goroutines-patterns/producer-consumer"

func main() {
	producerConsumer()

}

func producerConsumer() {
	execute(producerconsumer.ProducerConsumerPattern{})

	execute(producerconsumer.RealWorldPattern{})

	execute(producerconsumer.ProducerConsumerAdvanced{})
}

func execute(patterns Patterns) {
	patterns.Execute()
}
