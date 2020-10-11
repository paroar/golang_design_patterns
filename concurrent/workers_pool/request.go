package main

import (
	"fmt"
	"log"
	"sync"
)

type Request struct {
	Data    interface{}
	Handler RequestHandler
}

type RequestHandler func(interface{})

func NewStringRequest(s string, wg *sync.WaitGroup) *Request {
	myRequest := &Request{
		Data: s,
		Handler: func(i interface{}) {
			defer wg.Done()
			s, ok := i.(string)
			if !ok {
				log.Fatal("Invalid casting to string")
			}
			fmt.Println(s)
		},
	}
	return myRequest
}
