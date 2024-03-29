// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package greeting should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
// While documentation is taken very seriously in Go, and tools like golint will
// complain if you don't include it, these comments are not required to make
// the test suite pass. It is up to you whether to include them in your submitted
// exercises.
// You can read more about the importance of documentation in Go here:
// https://blog.golang.org/godoc-documenting-go-code
package greeting

// testVersion identifies the version of the test program that you are
// writing your code to. If the test program changes in the future --
// after you have posted this code to the Exercism site -- reviewers
// will see that your code can't necessarily be expected to pass the
// current test suite because it was written to an earlier test version.
//
// This is a convention done for Exercism exercises in the Go language track,
// it is not a requirement of the Go programming language.
//
// This test versioning setup will be common to all the exercises in the
// Go language track. When crafting your own solution file from scratch you
// will be expected to add this constant or the initial test will fail.
// The version number you should use will be found in the constant
// "targetTestVersion" in the test file, see ./hello_test.go for more
// information.
const testVersion = 4

func HelloWorld() string {
	return "Hello, World!"
}
