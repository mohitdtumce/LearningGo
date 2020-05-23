package main

import (
	"fmt"
	"math"
)

// Constant declared at package level.
const random string = "Mohit Sharma"

// Constant block
const (
	// iota is like a counter
	_ = iota + 5
	alpha
	beta
	gamma
)

// We are using a blank identifier.
// => Yes we know this is going to return a value but we don't care what value it returns.
const (
	_  = iota
	KB = float64(1 << (10 * iota))
	MB
	GB
	TB
	PB
	EB
	ZB
)

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancial

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	// ==================================================
	// 4.0 Constants

	// ==================================================
	// 4.1 Naming Convention
	// All constants are proceeded with the const keyword.
	// If you want to export the const to the outside world, use PascalCase
	const Exportable float64 = math.Pi
	fmt.Printf("%v, %T\n", Exportable, Exportable)
	// Else you can just use normal camelCase
	const myConst int = 101
	fmt.Printf("%v, %T\n", myConst, myConst)

	// Needless to say that reassigning value of constant (for ex: myConst = 32) won't work.
	// Value of constant must be determined at compile time and not runtime.
	// Therefore (const myAnotherConst float64 = math.sin(1.57)) won't work as well.
	// This is because value of math.sin(1.57) will be determined at runtime.

	// Constants can be made up of any primitive types that we discussed earlier.
	const a int = 12
	const b float64 = 3.14
	const c bool = true
	const d complex128 = 1 + 1i
	const e string = "This is also a constant"
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
	fmt.Printf("%v, %T\n", e, e)

	// Note -> Arrays will always be variable type
	// Note -> Constants can be shadowed. Notice we declare a constant shadow on package level.
	// We can redeclare the constant inside the block. Inner block will get precedence.
	fmt.Printf("%v, %T\n", random, random)
	const random float64 = 10000.0
	fmt.Printf("%v, %T\n", random, random)
	// ==================================================
	// 4.2 Typed constants and untyped constants
	// So far we've discussed typed constants where we specify type of the constant in the declaration statement.
	// But we don't necessarily have to do that. We can use compiler's ability to infer the type for us.
	const myYetAnotherConst = 3.14
	fmt.Printf("Untyped constant: %v, %T\n", myYetAnotherConst, myYetAnotherConst)

	const w = 42
	var x = 12.
	var y = 1 + 32i
	fmt.Printf("%v, %T\n", w+x, w+x) // Compiler sees this as 42 + x and interpret the type based on the type of x
	fmt.Printf("%v, %T\n", w+y, w+y) // Notice the type of w+x and w+y
	// So compiler did implicit type conversion in case of untyped constants which we cannot do in case of variables

	// ==================================================
	// 4.3 Enumerated Constants
	fmt.Printf("%v\n", alpha)
	fmt.Printf("%v\n", beta)
	fmt.Printf("%v\n", gamma)

	// ==================================================
	// 4.4 Enumeration Expression
	filesize := 4000000000.
	fmt.Println(KB, MB, GB, TB, PB, EB, ZB)
	fmt.Printf("%.2fGB\n", filesize/GB)

	var roles byte = isAdmin | canSeeFinancial | canSeeEurope
	fmt.Printf("%b\n", roles)

	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
}
