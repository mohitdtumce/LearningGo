package main

import (
	"fmt"
)

func main() {
	// ==================================================
	// 10.0 Pointers - 
	/*
		1. Creating Pointers, 
		2. Dereferencing pointers, 
		3. The new function, 
		4. Working with nil and 
		5. Types with internal pointers
	*/
	
	var a int = 42
	// b is a pointer to integer which is holding memory address of a
	var b *int = &a 
	
	// To get the address of a variable, we use & operator
	fmt.Println("Memory address where a is stored: ", &a, b)
	
	// Dereferencing the pointer will give us the value stored at the memory address which pointer is pointing to.
	// Type of the pointer will specify how many bytes to read at the pointed memory address.
	fmt.Println("Value stored at the location is: ", a, *b) 		
	
	a = 27
	fmt.Println("Value stored at the location is: ", a, *b)
	
	*b = 36 
	fmt.Println("Value stored at the location is: ", a, *b)
	
	// // Pointer arithmetics :- Go does not support pointer arithmetics to keep it simple.
	// If however you wish to use pointer arithmetics in the code, you can use "unsafe" package
	// https://golang.org/pkg/unsafe/
	p := [3]int32{1, 2,3 }
	q := &p[0]
	r := &p[1]
	fmt.Printf("%v, %p, %p, %p", p, &p, q, r)

	// Pointer to user defined object struct
	var mPtr *myStructs
	mPtr = &myStructs{foo: 42}
	fmt.Println(mPtr)

	// The new function
	var nPtr *myStructs
	fmt.Println(nPtr) 		// Will hold the value nil
	nPtr = new(myStructs)	// We got uninitialized object here
	(*nPtr).foo = 36		// Dereferencing the pointer to set the value of objects
	fmt.Println(*nPtr) 


	// Primitives, Arrays, String, Object are passed by value unless we are using pointers
	mohit := myStructs{foo: 42}
	rohit := mohit
	rohit.foo = 36
	fmt.Println(mohit, rohit)

	// Maps, Slices are passed by reference
	amohit := map[string]string {
		"foo":"bar",
		"baz":"buz",
	}
	arohit := amohit
	arohit["foo"] = "car"
	fmt.Println(amohit, arohit)

}

type myStructs struct {
	foo int
}