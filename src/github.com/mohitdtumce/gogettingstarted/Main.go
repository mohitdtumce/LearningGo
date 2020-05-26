package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// ==================================================
	// 9.0 Defer, Panic and Recovery

	/*
	A defer statement defers the execution of a function until the surrounding function returns
	*/
	fmt.Println("Start")
	defer fmt.Println("Middle-I")
	defer fmt.Println("Middle-II")
	fmt.Println("End")
	/*
	Deferred functions execute in LIFO order (Stack). So the output of the above snippet of code will be :-
	Start
	End
	Middle-II
	Middle-I
	*/

	// Example - 2
	// Open the resource
	res, err := http.Get("https://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Close the resource
	defer res.Body.Close()
	/* 
		If you are working with a lot of resources then you should not use the defer statement. 
		This is because all the opened resources will be closed only once the application executes 
		and in the meantime those opened resouces will clog the memory
	*/

	robots, err := ioutil.ReadAll(res.Body)
	
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)

	// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
	// So, the output of the below code snipet will be "start" and not "end"
	a := "start"
	defer fmt.Println(a)
	a = "end"


}
