package pipeline

import "fmt"

func RunPipeline() {
	nums := make(chan int)
	squares := make(chan int)
	final := make(chan int)
	go Generator(nums, 20)
	go Mutator(Squared, nums, squares)
	go Mutator(Adder(1000), squares, final)
	Printer(final)
}

func Generator(out chan<- int, max int) {
	for i := 0; i <= max; i++ {
		out <- i
	}
	close(out)
}

func Mutator(fn func(a int) int, in <-chan int, out chan<- int) {
	for num := range in {
		out <- fn(num)
	}
	close(out)
}

func Squared(num int) int {
	return num * num
}

func Adder(num int) func(a int) int {
	return func(a int) int {
		return num + a
	}
}

func Printer(in <-chan int) {
	for result := range in {
		fmt.Println(result)
	}
}
