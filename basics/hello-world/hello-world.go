// packages act as namespaces...need 1 per project, but can have multiple...main is entry point to app
package main

// imports external libraries
import (
	"fmt"
)

// needs to be called main (only 1 per project) for program startup
func main() {
	fmt.Print("Hello world!")
}

// go run <file>
// go mod init example.com/test-app
// go build
