package generator

import "fmt"

type ExamplePattern struct{}

func evenGenerator(max int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i <= max; i += 2 {
			out <- i
		}
		close(out)
	}()
	return out
}

func (g ExamplePattern) Execute() {
	for num := range evenGenerator(10) {
		fmt.Println(num)
	}
}
