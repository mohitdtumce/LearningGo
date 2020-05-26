package main

import (
	"fmt"
)

func main() {
	// ==================================================
	// 8.0 Looping

	// Simple for loop
	for i := 0; i < 5; i += 2 {
		fmt.Println(i)
	}

	// More than one iterators
	for i, j := 1, 0; i < 20; i, j = i*2, j+1 {
		fmt.Println(i, j)
	}

	// Now iterator is scoped to main method instead of just the for-loop
	// Iterator and Increment expressions inside the for-loop body
	itr := 0
	for itr < 5 {
		fmt.Println("Another form", itr)
		itr++
	}

	// break keyword
	jtr := 1
	for {
		fmt.Print(jtr, " ")
		if jtr == 10 {
			fmt.Println()
			break
		}
		jtr++
	}

	// continue keywork
	for i := 1; i < 4; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// Nested for-loops
	for i := 1; i <= 4; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// To break from all the loops under the label 'MyLabel'
	MyLabel:
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if i*j >= 8 {
				fmt.Println()
				break MyLabel
			}
			fmt.Print(i * j)
		}
		fmt.Println()
	}

	// Iterating over a collection using for-range
	s := []int{11, 12, 13}
	for k, v := range s {
		fmt.Println(k, v)
	}

	covidCases := map[string]float32{
		"USA":    1.708,
		"Brazil": 0.376,
		"Russia": 0.362,
	}

	for k, v := range covidCases {
		fmt.Println(k, v)
	}

	for k := range covidCases {
		fmt.Print(k, )
	}

	for _, v := range covidCases {
		fmt.Print(v, )
	}

	str := "Hello World!!"
	for k, v := range str {
		fmt.Println(k, string(v))
	}
}
