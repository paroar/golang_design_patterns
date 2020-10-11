package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	bufferSize := 100
	var d dispatcher = NewDispatcher(bufferSize)

	workers := 10
	for i := 0; i < workers; i++ {
		var w IWorkerLauncher = &PreffixSuffixWorker{
			prefixS: fmt.Sprintf("WorkerID: %d -> ", i),
			suffixS: " World",
			id:      i,
		}
		d.LaunchWorker(w)
	}

	requests := 100

	var wg sync.WaitGroup
	wg.Add(requests)

	for i := 0; i < requests; i++ {
		req := NewStringRequest(fmt.Sprintf("(Msg_id: %d) -> Hello", i), &wg)
		d.MakeRequest(*req)
	}
	d.Stop()
	wg.Wait()

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
