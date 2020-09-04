package main

import (
	"fmt"
	"io"
	"net/http"
)

type LogServer struct {
	Handler   http.Handler
	LogWriter io.Writer
}

func (l *LogServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(l.LogWriter, "Request URI: %s\n", r.RequestURI)
	fmt.Fprintf(l.LogWriter, "Host: %s\n", r.Host)
	fmt.Fprintf(l.LogWriter, "Content Length: %d\n", r.ContentLength)
	fmt.Fprintf(l.LogWriter, "Method: %s\n", r.Method)
	fmt.Fprint(l.LogWriter, "-----------------------------------------\n")

	l.Handler.ServeHTTP(w, r)
}
