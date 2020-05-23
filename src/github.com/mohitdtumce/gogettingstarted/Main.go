package main

import (
	"fmt"
	"math"
	"strconv"
)

// Variable declaration/initialization at package level
var fullname = "Mohit Sharma"

// Variable declaration/initialization block at package level
var (
	name = "Mohit"
	age int8 = 24
	company = "Microsoft"
	salary float32 = 1000
	increment = 10.0
	description string
)

func main()  {
	// ==================================================
	// 2.1 Variable Declaration
	
	// Method 1
	var i int
	i = 31

	// Method 2
	var j int = 32

	// Method 3
	k := 33

	fmt.Println(i, j, k)

	// We can define variables on package level as well. 
	// But keep in mind that automatic inference of variable types(as shown in method 3) won't work there
	
	fmt.Printf("%v %T\n", name, name)
	fmt.Printf("%v %T\n", age, age)
	fmt.Printf("%v %T\n", company, company)
	fmt.Printf("%v %T\n", salary, salary)
	fmt.Printf("%v %T\n", increment, increment)
	description = "Software Developer"
	fmt.Printf("%v %T\n", description, description)

	// ==================================================
	// 2.2 Re-declaration and shadowing

	// Note -> A variable must be used if declared
	// Note -> There are 3 scopes -> block, local, global
	// Re-declaration of variable 'fullname' which is also declared on package level outside main method.
	// Local variable declared within the block will take precendence => called Shadowing. 
	// Outer variable is still available but it is hidden.
	fmt.Println(fullname)
	var fullname string = "Rohit Saxena"
	fmt.Println(fullname)

	// ==================================================
	// 2.3 Visibility
	// Lower case variables are scoped to the packages. 
	// However if we use upper case variables then Go compiler will expose them to the outside world
	var Person string = "Mohit Sharma"
	fmt.Println(Person)

	// ==================================================
	// 2.4 Naming Convention

	// ==================================================
	// 2.5 Type Conversion

	var x float64 = math.Pi*10
	fmt.Printf("%v, %T\n", x, x)

	var y int
	y = int(x)
	fmt.Printf("%v, %T\n", y, y)

	var z string
	z = strconv.Itoa(y)
	fmt.Printf("%v, %T\n", z, z)

}