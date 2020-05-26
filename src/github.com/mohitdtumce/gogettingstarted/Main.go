package main

import (
	"fmt"
)

func main() {
	// ==================================================
	// 5.0 Arrays and Slices

	// ==================================================
	// 5.1 [Arrays] Creation & Built-in functions
	// Arrays are collection of items with same type.
	// Arrays have fixed size.

	// Declaration - 1/3
	arr1 := [3]string{"Mohit", "Rohit", "Vishal"}
	fmt.Println(arr1)

	// Declaration - 2/3
	arr2 := [...]string{"Mohit", "Rohit", "Vishal"}
	fmt.Println(arr2)

	// Declaration - 3/3
	var arr3 [5]string 
	arr3 = [...]string{"Mohit", "Rohit", "Vishal", "Rajat", "Danish"}
	fmt.Println(arr3)

	var grades [5]int
	grades = [5]int{77, 95, 95, 88, 80}
	// Arrays are passed by value. So copies refer to different underlying data
	anotherGrades := grades  
	anotherGrades[0] = 100
	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Another Grades: %v\n", anotherGrades)

	const row, col = 5, 3
	matrix := [row][col]int{{}}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			matrix[i][j] = i + j
		}
	}
	fmt.Println(matrix)

	var identityMatrix [3][3]int = [3][3] int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	fmt.Println(identityMatrix)

	// To find the length of array, we can use inbuilt function len()
	fmt.Println(len(grades))
	// len() method returns the size of array
	fmt.Println(len(matrix), len(matrix[0]))
	fmt.Println(len(identityMatrix), len(identityMatrix[0]))

	// ==================================================
	// 5.2 [Slices] Creation & Built-in functions
	// Slices are backed by Arrays. They don't have fixed size.
	slc := []int{1, 2, 3, 4, 5, 6}
	slc1 := slc // Slices are passed by reference
	slc1[0] = 100
	fmt.Println(slc, slc1)
	
	fmt.Printf("Size: %v, %v\n", len(slc), len(slc1))
	fmt.Printf("Capacity: %v %v\n", cap(slc), cap(slc1))

	// Slicing operation
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]
	c := a[3:]
	d := a[:6]
	e := a[3:6]
	fmt.Println(a ,"\n",b ,"\n",c ,"\n",d ,"\n", e)

	// As mentioned, slices are passed by reference
	// So change in any slice object will change the underlying data.
	a[5] = 42
	fmt.Println(a ,"\n",b ,"\n",c ,"\n",d ,"\n", e)

	// Using in-built method to create slices
	f := make([]int, 5, 100)
	fmt.Println(f, len(f), cap(f))

	g := []int{}
	fmt.Println(g, len(g), cap(g))

	g = append(g, 1)
	fmt.Println(g, len(g), cap(g))

	g = append(g, 2, 3, 4, 5)
	fmt.Println(g, len(g), cap(g))

	g = append(g, []int{6, 7, 8, 9, 10}...)
	fmt.Println(g, len(g), cap(g))

	h := []int{11, 12, 13, 14}
	g = append(g, h...)
	fmt.Println(g, len(g), cap(g))

	// Slices as stacks and queues
	// Removing the first element (FIFO)
	m := []int{1, 2, 3, 4, 5}
	for len(m) > 0 {
		fmt.Println(m)
		m = m[1:]
	}
	// Removing the last element (LIFO)
	n := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for len(n) > 0 {
		fmt.Println(n)
		n = n[:len(n)-1]
	}

	// Removing a subarray of elements from the middle of the slice
	// For example: Removing 4, 5, 6
	o := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	o = append(o[:3], o[6:]...)
	fmt.Println(o)

}
