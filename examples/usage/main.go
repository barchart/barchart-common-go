package main

import (
	"flag"

	"github.com/barchart/common-go/pkg/logger"
	"github.com/barchart/common-go/pkg/parameters"
	"github.com/barchart/common-go/pkg/usage"
)

var log = logger.Log

func main() {
	parameters.Add("TITLE", "Example", "The application title", false)

	myParam := parameters.Parse()

	usage.Initialize(myParam["TITLE"].(string), "This is the example application to show how to work with the usage package.")
	usage.AddParameters()
	usage.AddCommand("age", "Print your age", "age")
	usage.AddCommand("name", "Print your name", "name")
	usage.AddArgument("name", "Print hello <name>")
	usage.AddExample("go run main.go age 30")
	usage.AddExample("go run main.go name Jay")
	usage.AddExample("go run main.go v0.0.5")

	switch flag.Arg(0) {
	case "age":
		{
			if flag.NArg() < 2 {
				log.Fatal(usage.GetUsage())
			}

			log.Printf("Your age is: %v", flag.Arg(1))
		}
	case "name":
		{
			if flag.NArg() < 2 {
				log.Fatal(usage.GetUsage())
			}

			log.Printf("Your name is: %v", flag.Arg(1))
		}
	default:
		if flag.NArg() == 0 {
			log.Fatal(usage.GetUsage())
		}

		log.Printf("Hello %v!", flag.Arg(0))
	}
}
