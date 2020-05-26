package main

import (
	"fmt"
	"math"
)

func returnTrue() bool {
	fmt.Println("Always return true")
	return true
}

func main() {
	// ==================================================
	// 7.0 Control Flow

	// ==================================================
	// 7.1 Operators, If statement, If-else statement, If-else-if Statement
	if true {
		fmt.Println("This is an if statement block")
	}

	statePopulations := map[string]int{
		"Delhi":       11,
		"Maharashtra": 12,
		"Gujrat":      13,
		"Goa":         14,
	}
	if state, ok := statePopulations["Goa"]; ok {
		fmt.Println(state)
	}

	// If-else statement and If-else ladder
	number := 50
	guess := 30

	if guess < 1 || guess > 100 {
		fmt.Println("Number should be between 1-100")
	} else {
		if guess < number {
			fmt.Println("Too Low")
		} else if guess == number {
			fmt.Println("Correct Answer")
		} else {
			fmt.Println("Too High")
		}
	}

	// Short-circuiting
	age := 15
	if age > 12 || returnTrue() {
		// Since age > 12 is true, Go compiler will not evaluate rest of the expressions in ||
		fmt.Println("Passed Teen")
	}

	if age < 12 && returnTrue() {
		// Since age < 12 is false, Go compiler will not evaluate rest of the expressions in &&
		fmt.Println("Yet to be a teen")
	}

	myNum := 0.12345
	if myNum == math.Pow(math.Sqrt(myNum), 2) {
		fmt.Println("These are same")
	} else {
		fmt.Println("These are different")
	}
	// Answer comes out to be "These are different"
	// Instead use something like this :-
	if math.Abs(myNum - math.Pow(math.Sqrt(myNum), 2)) < 0.0001{
		fmt.Println("These are same")
	} else {
		fmt.Println("These are different")
	}

	// ==================================================
	// 7.2 Simple Cases, Multiple tests, Falling through, Type switches

	switch number := 16;number {
	case 3, 6, 9:
		fmt.Println("Multiple of 3")
	case 4, 8, 12:
		fmt.Println("Multiple of 4")
	case 5, 10, 15:
		fmt.Println("Multiple of 5")
	default:
		fmt.Println("Unhandled Case")
	}

	// Tag-less switch statement
	// Notice that break keywork is implied here.
	itr := 10
	switch {
	case itr <= 10:
		fmt.Println("Less than or equal to 10")
	case itr <= 20:
		fmt.Println("Less than or equal to 20")
	default:
		fmt.Println("Greater than 20")
	}

	// If however we want to go and check other cases as well, we have to use the keywork fallthrough
	name := "Mohit"
	switch {
	case name == "Mohit":
		fmt.Println("MySelf")
		fallthrough
	case name == "Mohit", name == "Rohit", name == "Vishal":
		fmt.Println("Friends")
	default:
		fmt.Println("Enemies")
	}

	// Interface tag can take any type assigned to it.
	// var jtr interface{} = 1
	// var jtr interface{} = []int{}
	// var jtr interface{} = "1"
	var jtr interface{} = 1.

	switch jtr.(type) {
	case int:
		fmt.Println("jtr is an int")
	case float64, float32:
		fmt.Println("jtr is a floating point number")
	case string:
		fmt.Println("jtr is a string")
	case []int:
		fmt.Println("jtr is a slice")
	default:
		fmt.Println("jtr has unknown type")
	}
}
