package main

import (
	"fmt"
	"time"

	s "github.com/leepuppychow/concurrency-practice/surfaces"
)

func main() {
	start := time.Now()

	fmt.Println(s.GeneratePlot())

	fmt.Println(time.Since(start))
}
