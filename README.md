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

#### Install Dependencies

Run ```go mod download``` to install dependencies.

#### Release New Code

* Add a new file for release notes to the ```./releases``` folder,
* Create a new tag, using the ```./tag.sh``` script, and
* Create a [GitHub Release](https://github.com/barchart/common-go/releases), using the aforementioned release notes.

#### Run Unit Tests

Run ```make test-v``` to execute unit tests.

### License

This software is provided under the MIT license.