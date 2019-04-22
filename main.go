package main

import (
	"flag"
	"sync"

	"github.com/Ahmah2009/myhttp/lib"
)

var wg sync.WaitGroup

func main() {
	parallelCount, addressList := readArguments()

	input := make(chan string, len(addressList))
	workersLimit := parallelCount

	// Increment waitgroup counter and create go routines
	for i := 0; i < workersLimit; i++ {
		wg.Add(1)
		go worker(input, i)
	}

	// Producer: load up input channel with jobs
	for _, job := range addressList {
		input <- job
	}

	// Close input channel since no more jobs are being sent to input channel
	close(input)
	// Wait for all goroutines to finish processing
	wg.Wait()

}

func worker(input chan string, id int) {
	// send done to decrement the wait groupe counter
	defer wg.Done()

	// fetch a task from the input chanel
	for value := range input {
		lib.GetURLResponce(value)
	}
}
func readArguments() (int, []string) {
	// get the command line argument flag value of "parallel"
	parallelPtr := flag.Int("parallel", 10, "")
	flag.Parse()

	return *parallelPtr, flag.Args()
}
