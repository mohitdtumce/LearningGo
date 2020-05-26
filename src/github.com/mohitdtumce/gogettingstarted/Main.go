package main

import (
	"fmt"
	"reflect"
)

const (
	deptIT int = 101
	deptHR int = 202
	deptPR int = 303
	deptEE int = 404
)

/*
Struct can be compared with classes in OOP
Struct is a collection of disparate data types that describe a single concept.
Since employee follows camelCase, it is local to the package and cannot be accessed outside the package.
If you wish to export the struct outside the package, simply use PascalCase.

You'll notice that attribute names are also written in camelcase.
It follows the general naming convention of Go as well
which means they cannot be accessed in other packages.
*/
type employee struct {
	id       int
	fullName string
	deptCode int
	projects []string
}

// Animal struct
type Animal struct {
	Name   string
	Origin string
}

// Bird struct
type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

// Doctor struct
type Doctor struct {
	Name       string `required max:100`
	Speciality string
}

func main() {
	// ==================================================
	// 6.0 Maps and Structs

	// ==================================================
	// 6.1 [Maps] Creation and Manipulation

	// Creation: Literal syntax to create a map
	statePopulations := map[string]int{
		"California": 12,
		"Texas":      13,
		"Florida":    14,
	}
	fmt.Println(statePopulations)

	/* Basic constraint on a key while creating a map is that it can be tested for equality.
	Most of the types can be used for key but not all. For example slices and another map.
	m := map[[]int]int {}  will throw an error saying "Invalid map key type"
	However arrays are fine. So below code will work just fine.
	*/
	m := map[[3]int]int{}
	fmt.Println(m)

	// Creation: Inbuilt make function to create map
	countryPopulations := make(map[string]int)
	countryPopulations = map[string]int{
		"India": 100,
		"US":    25,
		"China": 32,
	}
	fmt.Println(countryPopulations)

	// Fetching the value associated with the given key
	fmt.Println(countryPopulations["India"])   // Key present
	fmt.Println(countryPopulations["Britain"]) // Key not present will return the value 0

	// To check if a key is present in the map, we can use below written syntax as well.
	_, ok := countryPopulations["Britain"]
	fmt.Println(ok)
	_, ok = countryPopulations["India"]
	fmt.Println(ok)

	// Adding new key
	countryPopulations["Australia"] = 12
	fmt.Println(countryPopulations)
	fmt.Println(countryPopulations["Australia"])

	// Important: Order of keys is not guaranteed in maps.

	// Deleting keys from the map
	fmt.Println(countryPopulations)
	delete(countryPopulations, "Australia")
	fmt.Println(countryPopulations)

	// len() method will provide the length of map
	fmt.Println(len(countryPopulations))

	// Map is passed by reference
	anotherCountryPopulations := countryPopulations
	delete(anotherCountryPopulations, "India")
	fmt.Println(countryPopulations)
	fmt.Println(anotherCountryPopulations)

	// ==================================================
	// 6.2 [Structs] Creation, Naming Conventions, Embedding, and Tags

	counter := 0
	mohit := employee{
		id:       counter,
		fullName: "Mohit Sharma",
		deptCode: deptIT,
		projects: []string{
			"DC Migration",
			"Preorder and Partpayment",
		},
	}
	counter++

	rohit := employee{
		id:       counter,
		fullName: "Rohit Sharma",
		deptCode: deptEE,
		projects: []string{
			"Power Plant",
		},
	}
	counter++

	// Accessing specific attribute
	fmt.Println("Mohit's Projects: ", mohit.projects[0])
	fmt.Println("Fullname ", mohit.fullName)

	employees := []employee{}
	employees = append(employees, []employee{mohit, rohit}...)
	fmt.Println(employees)

	// Anonamous struct
	vishal := struct {
		name string
		age  int8
	}{
		name: "Vishal Singh",
		age:  42,
	}
	fmt.Printf("%v,  %T\n", vishal, vishal)

	vishal.age = 26
	fmt.Printf("%v,  %T\n", vishal, vishal)

	// Unlike maps and slices, structs are passed by values
	anotherVishal := vishal
	anotherVishal.age = 101
	fmt.Println(vishal, '\n', anotherVishal)

	// If you do want to point to the same data, we can use & (address of) operator.
	sameVishal := &vishal
	sameVishal.age = 87
	fmt.Println(vishal, '\n', sameVishal)

	// Embedding - Go doesn't have inheritence. It uses something called composition through embedding
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)

	c := Bird{
		Animal:   Animal{Name: "Penguin", Origin: "NewZealand"},
		SpeedKPH: 12,
		CanFly:   false,
	}

	fmt.Println(c)

	// Tags
	t := reflect.TypeOf(Doctor{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field)
	fmt.Println(field.Tag)
}
