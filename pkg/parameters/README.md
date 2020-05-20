# Parameters  
  
The `parameters` package provides a possibility to define and parse parameters from `flags`,   
`environment variables`, and the `AWS Secrets Manager` service.  

> `parameters.Parse()` will panic() if required parameters wasn't provided.

## Options

The Parameters package can be setup by providing `options` structure as last parameter. 

**Example**: 

```go
parameters.Add("EXAMPLE_SECRET", "", "An example parameter from AWS", true, parameters.Options{SecretsManagerEnable: true})
```

The options structure has following fields:

```go
type Options struct {
	SecretsManagerEnable bool
	StageSensitive       bool
}
``` 

* `SecretsManagerEnable` - Searches a parameter value inside AWS Secrets Manager.
* `StageSensitive` - Searches a parameter value inside AWS Secrets Manager with a prefix with a value of the `STAGE` parameter. 

Here is an example of searching a parameter value with the`StageSensitive` option: 

```go
parameters.Add("STAGE", "DEV", "A stage parameter.", true)
parameters.AddDatabase("EXAMPLE_DATABASE", database.Database{}, "An example database parameter from AWS", true, parameters.Options{SecretsManagerEnable: true, StageSensitive: true})
myParams := parameters.Parse()
```

```shell
go run main.go --STAGE=DEV
```

1. The `STAGE` value is DEV.
2. The parameters package will search for `EXAMPLE_DATABASE` flag and env.
3. The parameters package will search for `EXAMPLE_DATABASE_DEV` value inside AWS Secrets Manager.
4. The value of the `EXAMPLE_DATABASE` parameter can be found by the `EXAMPLE_DATABASE` key in both cases.

  
## Examples  
  
All examples are stored inside the `/examples/parameters` folder.  
  
### With AWS Secrets Manager
  
**Path** - [`examples/parameters/with-aws-secretsmanager/main.go`](/examples/parameters/with-aws-secretsmanager/main.go)  
  
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
  
#### How to run  

You can run the example by following commands:

> make run
>
> make run-flags
>
> make run-env

or

> go run main.go  
>
> go run main.go --STAGE=prod --HOST="some host" --PORT=1234 --DATABASE=database_name --LOCAL=true  
>
> STAGE=DEV HOST="some host" PORT=1234 DATABASE="database_name" LOCAL=true go run main.go
  
### Without AWS Secrets Manager
  
**Path** - [`examples/parameters/without-aws-secretsmanager/main.go`](/examples/parameters/without-aws-secretsmanager/main.go)  
  
  
It reads parameters from the `flags` and `environment variables` and prints them to the console.   
  
#### How to run  

You can run the example by following commands:

> make run
>
> make run-flags
>
> make run-env

or

> go run main.go  
>
> go run main.go --STAGE=prod --HOST="some host" --PORT=1234 --DATABASE=database_name --LOCAL=true 
>
> STAGE=DEV HOST="some host" PORT=1234 DATABASE="database_name" LOCAL=true go run main.go
