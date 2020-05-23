package main

import "fmt"

func main()  {
	// ==================================================
	// 3.0 Primitives

	// ==================================================
	// 3.1 Boolean
	var w bool
	var x bool = true
	var y bool = false
	z := (1 == 1)
	fmt.Printf("%v, %T\n", w, w)
	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", y, y)
	fmt.Printf("%v, %T\n", z, z)


	// ==================================================
	// 3.2 Numeric - Integer, Floating, Complex
	
	// Signed Integer Type 
	// int8 	[-128, 127]
	// int16	[-32 768, 32 767]
	// int32	[-2 147 483 648, 2 147 483 647]
	// int64	[-9 223 372 036 854 775 808, 9 223 372 036 854 775 807]

	// Unsigned Integer Type 
	// uint8 	[0, 255]
	// uint16	[0, 65 535]
	// uint32	[0, 4 294 967 295]
	// We don't have 64 bit unsigned integer

	// Note -> We cannot do operations on two operands of different types unless we do explicit type casting
	var i int8 = 10
	var j int = 2
	k := (int(i) + j)
	fmt.Printf("%v, %T\n", k, k)

	// Bitwise operators
	var a int = 10  		// 1010
	var b int = 3			// 0011
	fmt.Println(a & b)		// 0010 => 2
	fmt.Println(a | b)		// 1011	=> 11
	fmt.Println(a ^ b)		// 1001	=> 9

	var c int = 3 			// 0011
	fmt.Println(c << 1)		// 0110 => 6
	fmt.Println(c << 3)		// 11000 => 24
	fmt.Println(c >> 1)		// 0001 => 1
	fmt.Println(c >> 2)		// 0000 => 0

	// Floating Type
	// float32 [+/-1.8E(-38), +/- 3.4E(38)]
	// float64 [+/-2.23E(-308), +/- 1.8E(308)] default

	var m float32 = 3.14
	fmt.Printf("%v, %T\n", m, m)
	m = 13.7e-72
	fmt.Printf("%v, %T\n", m, m)
	m = 2.1e+14
	fmt.Printf("%v, %T\n", m, m)
	var n float64 = 13.7e-72
	fmt.Printf("%v, %T\n", n, n)
	
	p := 10.2
	q := 3.7
	fmt.Println(p + q)
	fmt.Println(p - q)
	fmt.Println(p * q)
	fmt.Println(p / q)

	// Complex Type
	// complex64 and complex128 (default)

	var r complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", real(r), real(r))
	var s = 2 + 2i
	fmt.Printf("%v, %T\n", s, s)

	// ==================================================
	// 3.3 Text

	str := "This is utf8 encoded => uint8."
	fmt.Printf("%v, %T\n", str, str)
	// Note -> String can be treated as collection of characters
	fmt.Printf("%v, %T\n", str[2], str[2])
	fmt.Printf("%v, %T\n", string(str[2]), str[2])
	// Note -> Strings are immutable. So str[2] = 'u' won't work.

	// String concatenation
	anotherStr := " This is also utf8 encoded."
	fmt.Printf("%v, %T\n", str + anotherStr, str + anotherStr)

	// Converting string into byte arrays
	byteArray := []byte(str)
	fmt.Printf("%v, %T\n", byteArray, byteArray)

	// rune represents utf32 int32 character
	var rn rune = 'a'
	fmt.Printf("%v, %T\n", rn, rn)
}