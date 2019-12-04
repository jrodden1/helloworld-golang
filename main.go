// This line declares the package that this file belongs to.
package main

/*
Typically 2 different types of packages:
Executable (generates a file that we can run), Reusable (helper logic/libraries)
'main' is a specific reserved word that indicates a executable will be made from this file
if any other word is used (like 'package cookies'), it will be a reusable / helper package

Also note, a package main file MUST have a func called main()
*/

import "fmt"

func main() {
	fmt.Println("Hello World!")
}

// To run this, go to the terminal then run:  go run main.go
// go run, compiles and immediately executes
// go build, compiles, does not execute, but creates executable file.  (see "main" file in this folder)
// go fmt, formats all the code in each file in the current directory
// go install and go get, manages packages
// go test runs any test files associated with the project.
