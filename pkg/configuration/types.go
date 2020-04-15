package configuration

import (
	. "github.com/barchart/common-go/pkg/configuration/aws"
	. "github.com/barchart/common-go/pkg/configuration/database"
)

// Databases is a slice of Database
type Databases map[string]Database

// Config is a type of configuration
type Config struct {
	Databases      Databases
	AWS            *AWS
	CustomSettings map[string]interface{}
	Stage          string
}
