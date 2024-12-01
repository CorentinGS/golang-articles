package generator

import "fmt"

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

type ExamplePattern struct{}

func (g ExamplePattern) Execute() {
	for num := range evenGenerator(10) {
		fmt.Println(num)
	}
}
