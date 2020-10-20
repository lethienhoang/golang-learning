package main

import (
	"fmt"
)

func main() {
	var x, y int
	// Initialize a waitgroup variable
	// wg := sync.WaitGroup{}
	// `Add(2) signifies that there is 2 task that we need to wait for
	// wg.Add(2)

	go func(x *int) {
		*x = 100
		// Calling `wg.Done` indicates that we are done with the task we are waiting
		// wg.Done()
	}(&x)

	go func(y *int) {
		*y = 10
		// Calling `wg.Done` indicates that we are done with the task we are waiting
		// wg.Done()
	}(&y)

	/*
	  Raca condition is when multiple threads are trying to access and manipulate the same variable.
	  the code below are all accessing and changing the value.
	  Divide x (100) by y (10) and the result is (10)...
	  Except if y is not assigned 10 before x,
	  y's initialized value of 0 is used,
	  You will see the runtime error: integer divide by zero in console log
	  Due to the uncertainty of Goroutine scheduling mechanism, the results of the following program is unpredictable,
	*/
	// wg.Wait()

	fmt.Println(x / y)
}
