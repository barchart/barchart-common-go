# @barchart/common-go

[![AWS CodeBuild](https://codebuild.us-east-1.amazonaws.com/badges?uuid=eyJlbmNyeXB0ZWREYXRhIjoiQnBnNGc5M3B3ZTlTMER2aHl6bEJuV1huQmJQdFFVdTMrMFJOMzVEMjU0MGR5VUZkNVVTcm54VVlpTUpNN2R3emg2SVoxNWsrc1BReE1zSmdZazZuN0l3PSIsIml2UGFyYW1ldGVyU3BlYyI6IkpCZEJOcVY1c2lYWW9XZTUiLCJtYXRlcmlhbFNldFNlcmlhbCI6MX0%3D&branch=master)](https://github.com/barchart/common-go)

A *public* library of shared [GoLang](https://golang.org/) utilities.
  
### Overview

#### Features 
  
* [Parameters](./pkg/parameters) - Pattern for accepting program arguments (e.g. flags, environment, AWS Secrets Manager).
* [Configuration](./pkg/configuration) - Pattern for storing configuration data (e.g. database connection).
* [Usage](./pkg/usage) - Utility to print usage (e.g. commands and arguments) for a program.
* [Logger](./pkg/logger) - Basic logging strategy.

### Development

#### Go Modules

Install dependencies as follows:

```shell
go mod download
```

#### Unit Tests

Execute unit test as follows:

```shell
make test-v
```

### License

This software is provided under the MIT license.