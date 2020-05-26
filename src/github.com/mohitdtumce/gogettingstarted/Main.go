package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// ==================================================
	// 9.0 Defer, Panic and Recovery

	// Classic example of panic situation
	a, b := 1, 0
	ans := a / b
	fmt.Println(ans)

	// Simple Webapplication which listens all the routes specified by "/" and port 8080
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Mushi Mushi"))
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		// For Instance if we try to run this application concurrently in different consoles, we get following error
		// panic: listen tcp :8080: bind: address already in use
		panic(err.Error())
	}

	// Example 3: defer executes before panic
	fmt.Println("Start")
	defer fmt.Println("This was deferred")
	panic("Something bad just happened")
	fmt.Println("End")
	/*
		Output of example 3 will be :-
		Start
		This was deferred
		panic: Something bad just happened
	*/

	// Example 4: Defer and Ananomous functions
	fmt.Println("Start")
	defer func ()  {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("Something bad has happened")
	fmt.Println("End")	// Unreachable code

	// Example 5
	fmt.Println("Start")
	panicker()
	fmt.Println("End")
	/*
		Output of example 5 :-
		Start
		About the panic
		2020/05/27 00:02:39 Error: Something bad has happened
		panic: Something bad has happened [recovered]
		        panic: Rethrowing panic if error cannot be handled
	*/
}

func panicker() {
	fmt.Println("About the panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			panic("Rethrowing panic if error cannot be handled")
		}
	}()
	panic("Something bad has happened")
	fmt.Println("Done Panicking")
}
