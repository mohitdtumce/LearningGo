package main

import (
	"fmt"
	"sync"
	"time"
)

// ==================================================
var wg = sync.WaitGroup{}

// ==================================================
const (
	logDebug   = "DEBUG"
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{})

func logger() {
	for {
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			break
		}

	}
}

func main() {
	// ==================================================
	// 14.0 Channels -
	/*
		1. Channel basics
		2. Restricting data flow
		3. Buffered channels
		4. Closing channels
		5. For...range loops with channels
		6. Select statement
	*/

	// Channels are the pipes that connect concurrent goroutines.
	// You can send values into channels from one goroutine and receive those values into another goroutine.

	// Channels are strongly typed => An integer channel can be used to send only int data
	ch := make(chan int)
	wg.Add(2)

	go func() {
		// Pulling data from the channel and assign it to the variable i.e in the direction of arrow
		i := <-ch
		fmt.Println("Pulling data from the channel", i)
		wg.Done()
	}()

	go func() {
		// data flows into the channel
		fmt.Println("Sending data into the channel")
		ch <- 42
		wg.Done()
	}()
	wg.Wait()

	// Example 2: Restricting data flow. Send only and receive only channel
	wg.Add(2)
	// This is receive only channel
	go func(ch <-chan int) {
		// Pulling data from the channel and assign it to the variable i.e in the direction of arrow
		i := <-ch
		fmt.Println("Pulling data from the channel", i)
		wg.Done()
	}(ch)

	// This is send only channel
	go func(ch chan<- int) {
		// data flows into the channel
		fmt.Println("Sending data into the channel")
		ch <- 42
		wg.Done()
	}(ch)

	wg.Wait()

	// Example 3: In this case there are 5 producer goroutines spawned inside for loop but just one consumer goroutine.
	// So, the following code snippet will give fatal error: all goroutines are asleep - deadlock!
	wg.Add(1)
	go func() {
		// Pulling data from the channel and assign it to the variable i.e in the direction of arrow
		i := <-ch
		fmt.Println("Pulling data from the channel", i)
		wg.Done()
	}()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			// data flows into the channel
			fmt.Println("Sending data into the channel")
			ch <- 42
			wg.Done()
		}()
	}
	wg.Wait()

	/*
		Now to solve this problen we use Buffered Channels.
		Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:
			ch := make(chan int, 100)

		By default channels are unbuffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value.
		Buffered channels accept a limited number of values without a corresponding receiver for those values.
	*/

	chnl := make(chan int, 5)
	wg.Add(2)
	go func(chnl <-chan int) {
		// For range construct for...range loop
		for {
			if val, ok := <-chnl; ok {
				fmt.Println(val)
			} else {
				break
			}
		}
		wg.Done()
	}(chnl)

	go func(chnl chan<- int) {
		chnl <- 12
		chnl <- 56
		// Closing a channel indicates that no more values will be sent on it.
		// This can be useful to communicate completion to the channel’s receivers.
		close(chnl)
		wg.Done()
	}(chnl)

	wg.Wait()

	// Select statement for channels
	go logger()
	logCh <- logEntry{
		time:     time.Now(),
		severity: logInfo,
		message:  "App is starting",
	}

	logCh <- logEntry{
		time:     time.Now(),
		severity: logInfo,
		message:  "App is shutting down",
	}

	time.Sleep(100 * time.Millisecond)
	doneCh <- struct{}{}
}
