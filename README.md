
# common-go  
  
Public shared library of utilities for Go (GoLang)  
  
## Libraries  
  
### Parameters  
  
The `parameters` package provides a possibility to define and parse parameters from `flags`,   
`environment variables`, and the `AWS Secrets Manager` service.  

> `parameters.Parse()` will panic() if required parameters wasn't provided.
  
### Examples  
  
All examples are stored inside the `/examples/parameters` folder.  
  
#### with-aws-secretsmanager
  
**Path** - [`examples/parameters/with-aws-secretsmanager/main.go`](examples/parameters/with-aws-secretsmanager/main.go)  
  
It reads parameters from the `flags`, `environment variables`, and the `AWS Secrets Manager` and prints them to the console.   
You should have access to the `AWS Secrets Manager` and the following secrets should exist:  
  
* **EXAMPLE_SECRET** - the plain text secret.   
* **EXAMPLE_SECRET_JSON** - the key-value pair.

**EXAMPLE_SECRET_JSON** have the following structure:
 
```json  
{ 
    "title": "some title",   
    "value": "some secret data"  
}  
```  
  
##### How to run  
  
**Examples:** 

> go run main.go  

> go run main.go --STAGE=prod --HOST="some host" --PORT=1234 --DATABASE=database_name --LOCAL=true  

> STAGE=prod go run main.go --HOST="some host" --PORT=1234 --DATABASE=database_name --LOCAL=true
  
#### without-aws-secretsmanager
  
**Path** - [`examples/parameters/without-aws-secretsmanager/main.go`](examples/parameters/without-aws-secretsmanager/main.go)  
  
  
It reads parameters from the `flags` and `environment variables` and prints them to the console.   
  
##### How to run  
  
**Examples:** 

> go run main.go  

> go run main.go --STAGE=prod --HOST="some host" --PORT=1234 --DATABASE=database_name --LOCAL=true 

> STAGE=prod go run main.go --HOST="some host" --PORT=1234 --DATABASE=database_name --LOCAL=true

### Configuration

The `configuration` package provides a possibility to store a configuration of 
`databases`, or AWS services such a `DynamoDB`, `SNS`, `SQS`, `SecretsManager` and e.t.c 

> Configuration package implements `Singleton` pattern. `configuration.GetConfig()` method
> returns the same instance of the configuration.

### Validation

The `validation` package is a wrapper over the `github.com/go-playground/validator`.  

> Validation package implements `Singleton` pattern. `validation.GetValidator()` method
> returns the same instance of the validator.Validate.