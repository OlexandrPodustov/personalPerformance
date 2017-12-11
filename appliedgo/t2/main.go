package main

import (
	"fmt"
)

func trace(name string) func() {
	fmt.Println("Entering ", name)
	// TODO:
	// 1. Print "Entering <name>"
	// 2. return a func() that prints "Leaving <name>"
	return func() {
		fmt.Println("Leaving ", name)
	}
}
func f() {
	ff := trace("aa")
	defer ff() // TODO: add trace() here so the defer receives the returned function

	fmt.Println("Doing something")
}
func main() {
	fmt.Println("Before f")
	f()
	fmt.Println("After f")
}

//https://appliedgo.com/courses/mastergo/lectures/2701459#/finished
//Task 2: Clever tracing with “defer”
//Intro
//The defer keyword allows to specify a function that is called whenever the current function ends. What if we could call one function at the beginning of the current function, and one at the end, with only one function call?
//
//Like so:
//
//func f() {
//	trace("f")
//	fmt.Println("Doing something")
//}
//And when calling function f() it would print:
//
//Entering f
//Doing something
//Leaving f
//With a tricky use of the defer statement and a closure, we can do that!
//
//Your task
//Write a function trace() that receives a string - the name of the current function - and does the following:
//
//Print “Entering <name>” where <name> is the string parameter that trace receives
//Create and return a function that prints “Leaving <name>”
//Then call trace() via the defer keyword in such a way that trace() runs immediately, and returns its result to defer.
//
//func trace(name string) func() {
//	// TODO:
//	// 1. Print "Entering <name>"
//	// 2. return a func() that prints "Leaving <name>"
//}
//func f() {
//	defer // TODO: add trace() here so the defer receives the returned function
//		fmt.Println("Doing something")
//}
//func main() {
//	fmt.Println("Before f")
//	f()
//	fmt.Println("After f")
//}
