package main

import (
	"fmt"
	"regexp"
	"sync"
	"testing"
)

func Test_Dispatcher(t *testing.T) {
	bufferSize := 100
	var d dispatcher = NewDispatcher(bufferSize)

	workers := 3
	for i := 0; i < workers; i++ {
		var w IWorkerLauncher = &PreffixSuffixWorker{
			prefixS: fmt.Sprintf("WorkerID: %d -> ", i),
			suffixS: " World",
			id:      i,
		}
		d.LaunchWorker(w)
	}

	requests := 10

	var wg sync.WaitGroup
	wg.Add(requests)

	for i := 0; i < requests; i++ {
		req := Request{
			Data: fmt.Sprintf("(Msg_id: %d) -> Hello", i),
			Handler: func(i interface{}) {
				defer wg.Done()
				s, ok := i.(string)
				if !ok {
					t.Fail()
				}
				ok, err := regexp.Match(`WorkerID\: \d* -\> \(MSG_ID: \d*\) -> [A-Z]*\sWorld`, []byte(s))
				if !ok || err != nil {
					t.Fail()
				}
			},
		}
		d.MakeRequest(req)
	}
	d.Stop()
	wg.Wait()
}
