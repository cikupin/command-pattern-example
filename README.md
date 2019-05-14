# Command Pattern Example
This is a simple examples of command pattern implementation in Go.

## Quick Start
### 1. Go and Dep Installation
##### A. Go
1. Download Go & follow the installation instruction found here: https://golang.org/doc/install

##### B. Go Dep
1. Go Dep is the official *"experiment"* Go package manager. Please follow the instruction here https://github.com/golang/dep

### 2. Install dependencies
Run this command from terminal:

    $ dep ensure

### 3. Run
To run:

    $ go run main.go

### 4. Run unit test
Use this command to check unit test coverage:

    $ go test $(go list ./... | grep -v /vendor/) -cover
