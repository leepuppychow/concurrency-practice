package pipeline

import "fmt"

func RunPipeline() {
	nums := make(chan int)
	squares := make(chan int)
	go generateNums(nums)
	go squareNums(nums, squares)
	printer(squares)
}

func generateNums(out chan<- int) {
	for i := 0; i <= 10; i++ {
		out <- i
	}
	close(out)
}

func squareNums(in <-chan int, out chan<- int) {
	for num := range in {
		out <- num * num
	}
	close(out)
}

func printer(in <-chan int) {
	for result := range in {
		fmt.Println(result)
	}
}
