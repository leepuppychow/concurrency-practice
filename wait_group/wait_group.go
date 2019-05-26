package wait_group

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func RunWaitGroup() {
	urls := []string{
		"https://www.google.com",
		"http://www.example.com",
		"https://docs.docker.com/",
		"https://www.geeksforgeeks.org/fundamentals-of-algorithms/",
		"https://www.google.com",
		"http://www.example.com",
		"https://docs.docker.com/",
		"https://www.geeksforgeeks.org/fundamentals-of-algorithms/",
		"https://www.google.com",
		"http://www.example.com",
		"https://docs.docker.com/",
		"https://www.geeksforgeeks.org/fundamentals-of-algorithms/",
		"https://www.google.com",
		"http://www.example.com",
		"https://docs.docker.com/",
		"https://www.geeksforgeeks.org/fundamentals-of-algorithms/",
	}
	getConcurrent(urls)
	getSync(urls)
}

func getConcurrent(urls []string) {
	fmt.Println("CONCURRENT - HTTP GETs")
	start := time.Now()
	var w sync.WaitGroup
	for _, url := range urls {
		w.Add(1)
		go func(url string) {
			defer w.Done()
			resp, _ := http.Get(url)
			fmt.Println(url, resp.StatusCode)
		}(url)
	}
	w.Wait()
	fmt.Println("\n\n\n", time.Since(start))
}

func getSync(urls []string) {
	fmt.Println("SEQUENTIAL - HTTP GETs")
	start := time.Now()
	for _, url := range urls {
		resp, _ := http.Get(url)
		fmt.Println(url, resp.StatusCode)
	}
	fmt.Println("\n\n\n", time.Since(start))
}
