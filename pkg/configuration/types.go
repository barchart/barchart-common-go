package configuration

import (
	. "github.com/barchart/barchart-common-go/pkg/configuration/aws"
	. "github.com/barchart/barchart-common-go/pkg/configuration/database"
)

type Databases map[string]Database

// Config is a type of configuration
type Config struct {
	Databases      Databases
	AWS            *AWS
	CustomSettings map[string]interface{}
	Stage          string
}
