package main

import (
	"encoding/json"
	"github.com/barchart/common-go/pkg/parameters"
	"log"
)

type exampleSecretJson struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

func main() {
	params := parameters.New()

	// Defining parameters
	// params.Add() is the alias to params.AddString()
	params.Add("STAGE", "dev", "", false)
	params.AddString("HOST", "", "", true)
	params.AddInt("PORT", 5432, "", false)
	params.Add("DATABASE", "", "", false)
	params.AddBool("LOCAL", false, "", true)

	// Defining parameters that will use AWS Secrets Manager.
	params.Add("EXAMPLE_SECRET", "", "", true, true)
	params.Add("EXAMPLE_SECRET_JSON", "", "", true, true)

	// This secret doesn't exist and the required attribute is true.
	// If the flag or environment variable wasn't provided, the value should be the default.
	params.Add("EXAMPLE_SECRET_DOEST_EXIST", "default value", "", false, true)

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

	log.Println("")
	log.Println("Unmarshal EXAMPLE_SECRET_JSON to the exampleSecretJson struct")
	log.Printf("EXAMPLE_SECRET_JSON: %+v", esj)
}
