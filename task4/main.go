package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	url := flag.String("url", "https://google.com", "Url address")
	reqAmount := flag.Int("reqAmount", 10, "Number of requests")
	timeout := flag.Int("timeout", 2500, "Timeout per request")
	flag.Parse()
	durations := executeRequests(*url, *reqAmount, *timeout)
	printStatistic(durations, *reqAmount)
}

func executeRequests(url string, reqAmount int, timeout int) (durations []int64) {
	mutex := &sync.Mutex{}
	var wg sync.WaitGroup
	for i := 0; i < reqAmount; i++ {
		wg.Add(1)
		go func(mutex *sync.Mutex, wg *sync.WaitGroup) {
			defer wg.Done()
			duration := handleGetExecution(url, timeout)
			mutex.Lock()
			durations = append(durations, duration)
			mutex.Unlock()
		}(mutex, &wg)
		wg.Wait()
	}
	return
}

func handleGetExecution(url string, timeout int) (duration int64) {
	startTime := time.Now()
	response, err := executeGet(url, timeout)
	duration = time.Since(startTime).Nanoseconds()
	if err != nil || response.StatusCode != 200 {
		duration = 0
	}
	return duration
}

func executeGet(url string, timeout int) (*http.Response, error) {
	client := http.Client{
		Timeout: time.Duration(timeout) * time.Millisecond,
	}
	return client.Get(url)
}

func printStatistic(durations []int64, reqAmount int) {
	var totalDuration int64
	minDuration := durations[0]
	maxDuration := durations[0]
	errorCounter := 0
	for _, duration := range durations {
		if duration != 0 {
			totalDuration += duration
			if minDuration > duration {
				minDuration = duration
			}
			if maxDuration < duration {
				maxDuration = duration
			}
		} else {
			errorCounter++
		}
	}
	successfulReqAmount := (int64)(reqAmount - errorCounter)
	var averageDuration int64
	if reqAmount != errorCounter {
		averageDuration = totalDuration / successfulReqAmount
	}
	fmt.Printf("Successfully execucted requests: %d/%d\n", successfulReqAmount, reqAmount)
	fmt.Println("Errors amount: ", errorCounter)
	fmt.Println("Total time for the requests: ", time.Duration(totalDuration)/time.Nanosecond)
	fmt.Println("Minimum time per request: ", time.Duration(minDuration)/time.Nanosecond)
	fmt.Println("Maximum time per request: ", time.Duration(maxDuration)/time.Nanosecond)
	fmt.Println("Average time per request: ", time.Duration(averageDuration)/time.Nanosecond)
}
