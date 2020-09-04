package main

import (
	"fmt"
	"net/http"
)

type BasicAuthMiddleware struct {
	Handler  http.Handler
	User     string
	Password string
}

func (b *BasicAuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()

	if ok {
		if user == b.User && pass == b.Password {
			b.Handler.ServeHTTP(w, r)
		} else {
			fmt.Fprintf(w, "User or password incorrect\n")
		}
	} else {
		fmt.Fprintln(w, "Error trying to retrieve data from Basic auth")
	}
}
