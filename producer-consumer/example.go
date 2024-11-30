package producer_consumer

import (
	"fmt"
)

type ProducerConsumerPattern struct{}

func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Printf("Produced: %d\n", i)
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for num := range ch {
		fmt.Printf("Consumed: %d\n", num)
	}
}

func (p ProducerConsumerPattern) Execute() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}
