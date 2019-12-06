// This line declares the package that this file belongs to.
package main

/*
Typically 2 different types of packages:
Executable (generates a file that we can run), Reusable (helper logic/libraries)
'main' is a specific reserved word that indicates a executable will be made from this file
if any other word is used (like 'package cookies'), it will be a reusable / helper package

Also note, a package main file MUST have a func called main()
*/

//this line allows us to import code from another package into ours.
import "fmt"

// 'fmt' specifically does printing to the terminal -- kind of like "puts" in ruby
// golang.org/pkg has a lot of documentation about built in / standard packages.

/*
short for Function - very similar to javascript / ruby methods etc.
func -> tells we are going to declare a function
main -> sets the name of the function
() -> list of args to pass to the function
{} -> inside the curlies is the code that will run when this fn is called.
*/
func main() {
	fmt.Println("Hello World!")
}

// To run this, go to the terminal then run:  go run main.go
// go run, compiles and immediately executes
// go build, compiles, does not execute, but creates executable file.  (see "main" file in this folder)
// go fmt, formats all the code in each file in the current directory
// go install and go get, manages packages
// go test runs any test files associated with the project.

// as for overall organization for a file.
// package main - Package Declaration (who / what am I)
// import fmt - import other packages that we need - from standard lib or from reusable packages
// declare functions, tell GO to do things
