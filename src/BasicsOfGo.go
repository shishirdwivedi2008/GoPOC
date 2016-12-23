package main

import "fmt"

func main() {
	fmt.Println("Hello my name is shishir")
	var x int
	x = 3
	y := 4
	sum, prod := addAndMultiply(x, y)
	fmt.Println("sum is ", sum, "product is:", prod)
	learnTypes()
	namedType1, namedType2 := learnNamedReturn(x, y)
	fmt.Println("first Return:", namedType1, "second return:", namedType2)
	learnLoops()
}

func addAndMultiply(x, y int) (sum, prod int) {
	return x + y, x * y
}

func learnTypes() {

	str := "Learn Go !"
	str1 := `A "raw" string literal
can include line breaks.`

	fmt.Println("Print String:", str, "Print Multi line string:", str1)

	f := 3.14554646457 //This is float value.
	complex := 3 + 4i  //This is complex number.

	fmt.Println("print float value", f, "Print complex number", complex)

	var arr [4]int //delcaration of array. At run time default value of int will be given //which is 0

	arr1 := [...]int{2, 4, 6, 78, 9, 90, 0}

	fmt.Println("arr with default value", arr, "array with init value", arr1)

}

func learnNamedReturn(x, y int) (c, d int) {
	c = x + y
	d = x * y
	return

}

func learnLoops() {
	x := 6
	if x > 5 {
		fmt.Println("value of x is greater that 5")
	}
	fmt.Println("Lets print table of 2")
	for i := 0; i <= 10; i++ {

		fmt.Println(i * 2)
	}

	//While loop.
	y := 0
	for {
		y++
		fmt.Println("Value of Y is :", y)
		if y > 100 {
			fmt.Println("Y value reached to 100")
			break
		}
	}
}
