package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestProCon(t *testing.T) {
	start := time.Now()
	stream := GetMockStream()

	// Producer
	tweets := producer(stream)

	var wg sync.WaitGroup
	wg.Add(2)
	// Consumer
	consumer(&wg, tweets, "worker1")
	consumer(&wg, tweets, "worker2")
	wg.Wait()
	fmt.Printf("Process took %s\n", time.Since(start))
}
