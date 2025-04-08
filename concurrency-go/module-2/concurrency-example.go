/*
Write two goroutines which have a race condition when executed concurrently.
Explain what the race condition is and how it can occur.
*/
package main

import "fmt"

var x int

func add_number() {
	x += 1
	// this bappens since the instruction is comprised of
	// more machine level instructions such as LOAD X, INCR and SAVE
	// Which leads to interlaving of two go routings with 3 instructions
	// which is non deterministic
}

func main() {
	x = 1
	go add_number()
	go add_number()
	fmt.Println(x)
}

/*
The above code has a race condition which is shown below in the comment

$ go run -race concurrency-example.go
0
==================
WARNING: DATA RACE
Write at 0x000005d6dd90 by goroutine 6:
  main.add_number()
      /Users/balart80/go/src/github.com/balart40/coursera-specialization-go/concurrency-go/module-2/concurrency-example.go:11 +0x3c

Previous read at 0x000005d6dd90 by main goroutine:
  main.main()
      /Users/balart80/go/src/github.com/balart40/coursera-specialization-go/concurrency-go/module-2/concurrency-example.go:18 +0x5c

Goroutine 6 (running) created at:
  main.main()
      /Users/balart80/go/src/github.com/balart40/coursera-specialization-go/concurrency-go/module-2/concurrency-example.go:16 +0x3e
==================
1
Found 1 data race(s)
exit status 66
*/
