package main

import (
	"fmt"
)

func sayMessage(msg, name *string) {
	fmt.Println(*msg, *name)
	*name = "Ted"
	fmt.Println(*name)
}

func sum(msg string, values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println(msg, result)
}

func summation(values ...int) int {
	fmt.Println("Summation returning result :-")
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

func summationPtr(values ...int) *int {
	fmt.Println("Summation returning pointer to local variable :-")
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}

func summationNamed(values ...int) (result int) {
	// Named returns
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	// Naked return
	return
}

func division(a, b float64) (float64, error) {
	fmt.Println("Multiple returns")
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

// Method for a struct
type greeter struct {
	greeting string
	name string
}

func (g greeter) sayHello() {
	g.name = "Rohit"
	fmt.Println(g.greeting + " " + g.name)
}

func (g *greeter) sayHi(){
	(*g).name = "Rohit"
	fmt.Println((*g).greeting + " " + (*g).name)
}

// /*
// Go does not support function overloading and does not support user defined operators.
// While Go still does not have overloaded functions (and probably never will), the most useful feature of overloading,
// that of calling a function with optional arguments and inferring defaults for those omitted can be simulated using
// a variadic function, which has since been added. But this comes at the loss of type checking
// */
// func sayMessage(msg, name string) {
// }

func main() {
	// ==================================================
	// 11.0 Functions -
	/*
		1. Basic Syntax,
		2. Parameters,
		3. Return values,
		4. Anonymous functions
		5. Functions as types
		6. Methods
	*/

	greetings := "Hello"
	name := "Stacey"
	// Apart from slices and maps which use internal pointers, variables are passed by values.
	// So if you want to make changes in the original copy, make sure to pass variables
	sayMessage(&greetings, &name)
	fmt.Println(name)

	// Variadic Parameters
	sum("The sum is:", 1, 2, 3, 4, 5)
	fmt.Println(summation(1, 2, 3, 4, 5))

	resultPtr := summationPtr(1, 2, 3, 4, 5)
	fmt.Println(*resultPtr)

	result := summationNamed(1, 2, 3, 4, 5)
	fmt.Println(result)

	quo, err := division(12.0, 0.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(quo)
	}

	// Anonymous function :- Generally used to generate isolated scope
	func ()  {
		fmt.Println("Anonymous Function")
	}()

	// Passing arguments in anomymous function
	func (name string, age int)  {
		fmt.Println(name, age)
	}("Mohit", 24)

	// Functions as types
	var f func() = func ()  {
		fmt.Println("Yo Bro!!")
	}

	f()

	divide := func(a, b float64) (float64, error){
		fmt.Println("Multiple returns")
		if b == 0.0 {
			return 0.0, fmt.Errorf("Cannot divide by zero")
		}
		return a / b, nil
	}

	d, e := divide(12.0, 5.0)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(d)
	}

	g := greeter { greeting: "Hello", name: "Mohit",}
	g.sayHello()
	fmt.Println(g.name)
	(&g).sayHi()
	fmt.Println(g.name)
}
