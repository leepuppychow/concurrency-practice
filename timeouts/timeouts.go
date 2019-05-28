package timeouts

import (
	"fmt"
	"net/http"
	"time"
)

func RunTimeoutExample(numRequests int) {
	done := make(chan time.Duration)
	go makeRequests(done, numRequests)

	select {
	case duration := <-done:
		fmt.Printf("%d requests performed in: %s\n", numRequests, duration)
	case <-time.After(1 * time.Second):
		fmt.Println("TIMED OUT AFTER 1 SECOND")
	}
}

func makeRequests(done chan time.Duration, numRequests int) {
	start := time.Now()
	for i := 0; i <= numRequests; i++ {
		http.Get("https://www.google.com")
	}
	done <- time.Since(start)
}
