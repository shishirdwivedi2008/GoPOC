package main

import "fmt"
import "math/rand"
import "time"

func isFunction() (x string) {
	return "Yes i am a function"
}

func parameterizedFunction(x, y int) {
	z := x + y
	fmt.Println("I don't return any value, i calculate and print", z)
}

func functionWithReturnValue(x, y string) (z string) {
	z = x + y
	return //Named return value automatically returns value of z
}

func funcWithMultipleReturnValue(x, y string) (a, b, c string) {
	a = x
	b = y
	c = x + y
	return a, b, c
}

func doSameRandomAddition() (z int) {
	var arr [100]int
	for x := 0; x < 100; x++ {
		rand.Seed(time.Now().UnixNano())
		y := rand.Intn(1000000)
		fmt.Println("Random number generated is", y)
		arr[x] = y
	}
	for x := 0; x < 100; x++ {
		z = z + arr[x]
	}
	return z
}

func functionOverloading(x int, y string) {
	fmt.Println(x, y)
}

//Function can return function

func myfunc(x string) func(y, z string) string {

	return func(y, z string) string {
		return fmt.Sprintf("%s %s %s", y, z, x)
	}
}

// Deferred statements are executed just before the function returns.

func learnDefer() (ok bool) {

	defer fmt.Println("deferred statements execute in reverse (LIFO) order.")
	defer fmt.Println("\nThis line is being printed first because")
	// Defer is commonly used to close a file, so the function closing the
	// file stays close to the function opening the file.
	return true
}

func main() {
	fmt.Println(isFunction())
	x := 10
	y := 20
	a := "Hello My Name is"
	b := "Shishir"
	parameterizedFunction(x, y)
	z := functionWithReturnValue(a, b)
	fmt.Println(z)
	a, b, c := funcWithMultipleReturnValue(a, b)
	fmt.Println("value of 1st Return", a, "value of 2nd return", b, "Value of 3rd return", c)
	functionOverloading(x, a)
	f := doSameRandomAddition()
	fmt.Println("Sum of array is :", f)
	fmt.Println(myfunc("shishir")("Hello", "My name is"))
	learnDefer()
}
