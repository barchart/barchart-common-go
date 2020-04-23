package main

import (
	"github.com/barchart/common-go/pkg/logger"
	"github.com/barchart/common-go/pkg/parameters"
)

var log = logger.Log

func main() {
	// Defining parameters
	// params.Add() is the alias to params.AddString()
	parameters.Add("STAGE", "dev", "", false)
	parameters.AddString("HOST", "", "", true)
	parameters.AddInt("PORT", 5432, "", false)
	parameters.Add("DATABASE", "", "", false)
	parameters.AddBool("LOCAL", false, "", true)

	// Parse all parameters. Will get value from flags, environment variables and AWS Secrets Manager.
	// parameters.Parse executes flag.Parse() under the hood. DON'T USE flag.parse()
	// parameters.Parse() returns map[string]interface{}
	myParams := parameters.Parse()

	// Work with parameters. Remember all values are interface{}. You should use a type assertion
	var local bool
	local = myParams["LOCAL"].(bool)
	log.Println("Assigns a parameter to a variable")
	log.Printf("LOCAL: %v \n", local)
	log.Println("")

	log.Println("Reads all parameter using for ... range:")
	log.Println("_______________")
	for k, value := range myParams {
		log.Printf("%v: %+v", k, value)
	}
	log.Println("_______________")
}
