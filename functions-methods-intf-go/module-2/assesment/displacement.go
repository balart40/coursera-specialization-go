package main

import "fmt"

/*
Let us assume the following formula for
displacement s as a function of time t, acceleration a, initial velocity vo,
and initial displacement so.

s = ½ a t2 + vot + so

Write a program which first prompts the user
to enter values for acceleration, initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the
program should compute the displacement after the entered time.

You will need to define and use a function
called GenDisplaceFn() which takes three float64
arguments, acceleration a, initial velocity vo, and initial
displacement so. GenDisplaceFn()
should return a function which computes displacement as a function of time,
assuming the given values acceleration, initial velocity, and initial
displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one
float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume
the following values for acceleration, initial velocity, and initial
displacement: a = 10, vo = 2, so = 1. I can use the
following statement to call GenDisplaceFn() to
generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to
print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print
the displacement after 5 seconds.

fmt.Println(fn(5))

Submit your Go program source code.
*/

func GenDisplaceFn(a float64, vo float64, so float64) func(t float64) float64 {
	fn := func(t float64) float64 {
		return 0.5*a*t*t + vo*t + so
	}
	return fn
}

func getFloat64(varName string) float64 {
	var value float64
	fmt.Printf("Provide %s: ", varName)
	_, err := fmt.Scan(&value)
	if err != nil {
		fmt.Printf("Error scanning %s input", varName)
	}
	return value
}

func main() {
	var a float64
	var vo float64
	var so float64
	var t float64
	// request the variables
	a = getFloat64("acceleration")
	vo = getFloat64(" initial velocity")
	so = getFloat64("initial displacement")
	t = getFloat64("time")
	// get function
	fn := GenDisplaceFn(a, vo, so)
	// print displacement in t
	fmt.Println(fn(t))
}
