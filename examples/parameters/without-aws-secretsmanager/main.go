package main

import (
	"github.com/barchart/common-go/pkg/logger"
	"github.com/barchart/common-go/pkg/parameters"
)

// go run main.go --STAGE=DEV
//
// go run main.go --STAGE=DEV --HOST="some host" --PORT=1234 --DATABASE=database_name --LOCAL=true
//
// STAGE=DEV go run main.go --HOST="some host" --PORT=1234 --DATABASE=database_name --LOCAL=true

var log = logger.Log

func main() {
	// Defining parameters
	// params.Add() is the alias to params.AddString()
	parameters.Add("STAGE", "DEV", "A stage parameter.", true)
	parameters.AddString("HOST", "", "A host of database", false)
	parameters.AddInt("PORT", 5432, "A port of database", false)
	parameters.Add("DATABASE", "", "A name of database", false)
	parameters.AddBool("LOCAL", false, "Run application locally", false)

	// Parse all parameters. Will get value from flags, environment variables and AWS Secrets Manager.
	// parameters.Parse executes flag.Parse() under the hood. DON'T USE flag.parse()
	// parameters.Parse() returns map[string]interface{}
	myParams := parameters.Parse()

	// Work with parameters. Remember all values are interface{}. You should use a type assertion
	var local bool
	local = myParams["LOCAL"].(bool)
	log.Println("Assigns a parameter to a variable")
	log.Printf("LOCAL: %v \n", local)

	log.Println("Reads all parameter using for ... range:")
	log.Println("_______________")
	for k, value := range myParams {
		log.Printf("%v: %+v", k, value)
	}
	log.Println("_______________")
}
