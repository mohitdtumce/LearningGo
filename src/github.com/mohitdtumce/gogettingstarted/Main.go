package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// ==================================================
// Example 1
func sayHello() {
	fmt.Println("Hello")
}

// ==================================================
// Example 3 : Using WaitGroups
var wg = sync.WaitGroup{}

var counter = 0

func greet() {
	fmt.Printf("Hola Mi Amigo #%v\n", counter)
	wg.Done()
}

func increment() {
	counter++
	wg.Done()
}

// ==================================================
// Example 4 : Using Mutex
var wgroup = sync.WaitGroup{}
var ctr = 0

/* 
	A RWMutex is a reader/writer mutual exclusion lock. 
	The lock can be held by an arbitrary number of readers or a single writer. 
	The zero value for a RWMutex is an unlocked mutex.
*/
var m = sync.RWMutex{}

func sayHi() {
	fmt.Printf("Hola Mi Amigo #%v\n", ctr)
	m.RUnlock()
	wgroup.Done()
}

func increase() {
	ctr++
	fmt.Println("Increasing ctr", ctr)
	m.Unlock()
	wgroup.Done()
}

func main() {
	// ==================================================
	// 13.0 Goroutines -
	/*
		1. Create Goroutines,
		2. Synchronization,
			a. WaitGroups
			b. Mutexes
		3. Parallelism
		4. Best Practices & Tools
	*/

	// Example 1: Simple Example of Goroutine
	go sayHello()
	time.Sleep(100 * time.Millisecond)

	// Example 2: Using Anonymous function
	/*
		Notice that we have created a dependency between main function and the goroutine.
		Our anonymous function here is dependent on main function for the variable msg. This can cause problem.
		Output for the below code snippet is "Goodnight". This is because in most of the cases compiler will not interrupt main() function unless it executes Sleep() method.
		It essentially means that even though it spawns another goroutine inside the main() function, it will not provide any CPU to newly spawned goroutine and it is still executing main() function.
		So instead of printing "Hello Again", it is printing "Goodnight". This is called race-condition and something which we want to avoid.
	*/
	var msg = "Hello Again"
	go func() {
		fmt.Println(msg)
	}()
	msg = "Goodnight"
	time.Sleep(100 * time.Millisecond)

	/*
		To solve the above mentioned problem, use parameters.
		Since string is passed by value, we have actually decoupled anotherMessage variable in the main() from the msg variable of the anonymous function.
	*/
	var anotherMessage = "Hello Yet Again"
	go func (msg string)  {
		fmt.Println(msg)
	}(anotherMessage)

	anotherMessage = "Goodnight Again"
	time.Sleep(100 * time.Millisecond)

	/*
		The solve mentioned above is not the best practices because we are using the Sleep call and thus binding/degrading the performance of the application.
		Alternative for that is to use WaitGroup. The task of WaitGroup is to synchronize multiple goroutines together.

		For ex: In this case we want to synchronize, we want to
	*/

	// Example 3: Using WaitGroups
	var message = "Hello"
	wg.Add(1)
	go func (message string)  {
		fmt.Println(message)
		wg.Done()
	}(message)
	message = "Adiós"
	wg.Wait()

	/*
		In the above example, we are synchronizing two goroutines together but only one of the goroutine is really doing the work.
		main() function is just storing data and spawing other goroutines.
		But we can have more than one goroutines working on the same data and we might need to synchronize those together.
	*/
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go greet()
		go increment()
	}
	wg.Wait()

	/*
		In the above example goroutines are racing against each other and there is not synchronization whatsoever.
		In order to correct this we need to find a way to synchronize these.
		We are going to introduce something called mutex.
	*/

	/*
		GOMAXPROCS sets the maximum number of CPUs that can be executing simultaneously and returns the previous setting.
		If n < 1, it does not change the current setting. The number of logical CPUs on the local machine can be queried with NumCPU.
		This call will go away when the scheduler improves.
	*/

	// Example 4: Locking resouces using Mutex
	runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		// main function is executing the locks and the spawned goroutines are asynchronously unlocking it once they're done.
		wgroup.Add(2)
		m.RLock()
		go sayHi()
		m.Lock()
		go increase()
	}
	wgroup.Wait()

	// But above process doesn't perform well because we are constantly locking/unlocking it. Infact it would have worked better without goroutines. However it can be beneficial in cases.

	fmt.Println(runtime.GOMAXPROCS(-1))

	// Best Practices
	/*
		1. Don't create goroutines in libraries. Let consumer control concurrency.
		2. When creating a goroutine, know how it will end. Aboid subtle memory leaks
		3. Check for race-condition at compile time. try running "go run -race /path/to/main/file/main_file.go"
		4.  
	*/
}
