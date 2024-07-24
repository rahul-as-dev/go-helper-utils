package main

import (
	"fmt"
)

/*
closure is a function defined inside another function. It has access to its outer function(s) scope.
variables declared in the function won't cease to exist once the control is returned form that function if
any inner function references them
*/
func main() {
	done := make(chan bool)

	// define a closure and immediately spawn it as a goroutine. Pass the done channel
	for i := 0; i < 10; i++ {
		go func(done chan bool) {
			fmt.Println("Go routine #", i) // with go < 1.22, all of this would print i as 10.
		}(done)
	}
	// wait for the goroutines to finish
	for i := 0; i < 10; i++ {
		<-done
	}
	//time.Sleep(1 * time.Second)
	fmt.Println("Done.")
}
