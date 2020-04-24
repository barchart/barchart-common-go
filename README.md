# @barchart/common-go

[![AWS CodeBuild](https://codebuild.us-east-1.amazonaws.com/badges?uuid=eyJlbmNyeXB0ZWREYXRhIjoiQnBnNGc5M3B3ZTlTMER2aHl6bEJuV1huQmJQdFFVdTMrMFJOMzVEMjU0MGR5VUZkNVVTcm54VVlpTUpNN2R3emg2SVoxNWsrc1BReE1zSmdZazZuN0l3PSIsIml2UGFyYW1ldGVyU3BlYyI6IkpCZEJOcVY1c2lYWW9XZTUiLCJtYXRlcmlhbFNldFNlcmlhbCI6MX0%3D&branch=master)](https://github.com/barchart/common-go)

A *public* library of shared GoLang utilities.
  
### Overview

#### Features 
  
* [Parameters](./pkg/parameters) - Pattern for accepting program arguments (e.g. flags, environment, AWS Secrets Manager)
* [Configuration](./pkg/configuration) - Pattern for storing configuration data (e.g. database connection)
* [Validation](./pkg/validation)
* [Usage](./pkg/usage)
* [Logger](./pkg/logger)

### Development

#### Go Modules

```sh
go get github.com/barchart/common-go
```

#### License

This software is provided under the MIT license.