package main

import (
	"flag"
	"log"
	"os"

	"github.com/paroar/design_patterns/behavioral/strategy/shapes"
)

func main() {
	var output = flag.String("output", "text", "The output to use between 'text' and 'image' file")

	flag.Parse()
	activeStrategy, err := shapes.FactoryPrinter(*output)
	if err != nil {
		log.Fatal(err)
	}
	switch *output {
	case shapes.TEXT_STRATEGY:
		activeStrategy.SetWriter(os.Stdout)
	case shapes.IMAGE_STRATEGY:
		w, err := os.Create("/tmp/image.jpg")
		if err != nil {
			log.Fatal("Error opening image")
		}
		defer w.Close()
		activeStrategy.SetWriter(w)
	}
	err = activeStrategy.Print()
	if err != nil {
		log.Fatal(err)
	}
}
