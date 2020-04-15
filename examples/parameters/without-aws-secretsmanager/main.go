package main

import (
	"github.com/barchart/common-go/pkg/parameters"
	"log"
)

func main() {
	params := parameters.New()

	// Defining parameters
	// params.Add() is the alias to params.AddString()
	params.Add("STAGE", "dev", "", false)
	params.AddString("HOST", "", "", true)
	params.AddInt("PORT", 5432, "", false)
	params.Add("DATABASE", "", "", false)
	params.AddBool("LOCAL", false, "", true)

	// Parse all parameters. Will get value from flags, environment variables and AWS Secrets Manager.
	// params.Parse executes flag.Parse() under the hood. DON'T USE flag.parse()
	// params.Parse() returns map[string]interface{}
	myParams := params.Parse()

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
