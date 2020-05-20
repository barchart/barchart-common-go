package main

import (
	"encoding/json"

	"github.com/barchart/common-go/pkg/configuration/database"
	"github.com/barchart/common-go/pkg/logger"
	"github.com/barchart/common-go/pkg/parameters"
)

// go run main.go --STAGE=DEV
//
// go run main.go --STAGE=DEV --HOST="some host" --PORT=1234 --DATABASE=database_name --LOCAL=true
//
// STAGE=DEV go run main.go --HOST="some host" --PORT=1234 --DATABASE=database_name --LOCAL=true

type exampleSecretJson struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

var log = logger.Log

func main() {
	// Defining parameters
	// params.Add() is the alias to params.AddString()
	parameters.Add("STAGE", "DEV", "A stage parameter.", true)
	parameters.AddString("HOST", "", "A host of database", false)
	parameters.AddInt("PORT", 5432, "A port of database", false)
	parameters.Add("DATABASE", "", "A name of database", false)
	parameters.AddBool("LOCAL", false, "Run application locally", false)

	// This parameter has StageSensitive: true, so in AWS Secrets Manager must have the following name: EXAMPLE_DATABASE_STAGE; Where STAGE is a value of the STAGE parameter.
	// e.g. EXAMPLE_DATABASE_DEV
	// To use a flag or an environment variable, the name should be the same as 1 argument: EXAMPLE_DATABASE
	// e.g. --EXAMPLE_DATABASE="some json string"
	// e.g. EXAMPLE_DATABASE="some json string" go run main.go
	// To get parameter from result of parse() function, the key should be the same as 1 argument: EXAMPLE_DATABASE
	parameters.AddDatabase("EXAMPLE_DATABASE", database.Database{}, "An example database parameter from AWS", true, parameters.Options{SecretsManagerEnable: true, StageSensitive: true})

	// Defining parameters that will use AWS Secrets Manager.
	parameters.Add("EXAMPLE_SECRET", "", "An example parameter from AWS", true, parameters.Options{SecretsManagerEnable: true})

	// This parameter should be stored as key/value secret in AWS Secrets Manager. The result of parsing will have a JSON string representing this parameter.
	parameters.Add("EXAMPLE_SECRET_JSON", "", "An example key/value parameter from AWS", true, parameters.Options{SecretsManagerEnable: true})

	// If this parameter wan't provided and doesn't exist in AWS Secrets Manager
	// the default value should be in the result of parse() function.
	parameters.Add("EXAMPLE_SECRET_DOES_NOT_EXIST", "default value", "", false, parameters.Options{SecretsManagerEnable: true})

	// Parse all parameters. Will get value from flags, environment variables and AWS Secrets Manager.
	// parameters.Parse executes flag.Parse() under the hood. DON'T USE flag.parse()
	// parameters.Parse() returns map[string]interface{}
	myParams := parameters.Parse()

	// Work with parameters. Remember all values are interface{}. You should use a type assertion
	var local bool
	log.Println("Assigns a parameter to a variable")
	local, isLocalProvided := myParams["LOCAL"].(bool)
	if isLocalProvided == true {
		log.Printf("LOCAL: %v \n", local)
	}

	log.Println("Reads all parameter using for ... range")
	log.Println("_______________")
	for k, value := range myParams {
		log.Printf("%v: %+v", k, value)
	}
	log.Println("_______________")

	// EXAMPLE_SECRET_JSON got as string from the AWS Secrets Manager, but it's a JSON (key-value pair) inside the string.
	// Here is an example of parsing this string to the structure.
	esj := exampleSecretJson{}
	_ = json.Unmarshal([]byte(myParams["EXAMPLE_SECRET_JSON"].(string)), &esj)

	log.Println("Unmarshal EXAMPLE_SECRET_JSON to the exampleSecretJson struct")
	log.Printf("EXAMPLE_SECRET_JSON: %+v", esj)
}
