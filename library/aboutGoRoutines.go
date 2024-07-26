package library

import "fmt"

func AboutGoRoutines() {
	/*
		GoLang Runtime has its own scheduler for implementing concurrency
		The runtime initially captures kernel threads equals to the number of cpu cores
		Each time a concurrent task is request, the goRoutine is pushed into a runtime-queue (runQueue)
		Each thread has its own runQueue
		There is also a global runQueue for the I/O tasks.
	*/
	done := make(chan bool)

	// define a closure and immediately spawn it as a goroutine. Pass the done channel
	for i := 0; i < 10; i++ {
		go func(i int, done chan bool) {
			fmt.Println("Go routine #", i) // with go < 1.22, all of this won't print i as 10 because we are passing a copy of the outer reference.
		}(i, done)
	}
	// wait for the goroutines to finish
	for i := 0; i < 10; i++ {
		<-done
	}
	//time.Sleep(1 * time.Second)
	fmt.Println("Done.")

}
