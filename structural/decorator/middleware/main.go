package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Enter the type number of server you want to launch from thefollowing:")
	fmt.Println("1.- Plain server")
	fmt.Println("2.- Server with logging")
	fmt.Println("3.- Server with logging and authentication")
	var selection int
	fmt.Fscanf(os.Stdin, "%d", &selection)

	switch selection {
	case 1:
		mySuperServer := new(MyServer)
		http.Handle("/", mySuperServer)
	case 2:
		mySuperServer := &LogServer{
			Handler:   new(MyServer),
			LogWriter: os.Stdout,
		}
		http.Handle("/", mySuperServer)
	case 3:
		var user, password string
		fmt.Println("Enter user and password separated by a space")
		fmt.Fscanf(os.Stdin, "%s %s", &user, &password)
		mySuperServer := &LogServer{
			Handler: &BasicAuthMiddleware{
				Handler:  new(MyServer),
				User:     user,
				Password: password,
			},
			LogWriter: os.Stdout,
		}
		http.Handle("/", mySuperServer)
	default:
		mySuperServer := new(MyServer)
		http.Handle("/", mySuperServer)
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}
