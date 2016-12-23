// MethodRecieverAndInterfaces
package main

import (
	"fmt"
)

//While technically Go isn’t an Object Oriented Programming language,
// types and methods allow for an object-oriented style of programming.
//The big difference is that Go does not support type inheritance but instead has
//a concept of interface.

type user struct {
	name, lastname string
}

func (u user) greetMe() {
	fmt.Println("My name is :", u.name, u.lastname)
}

type Person struct {
	name user
}

func (a Person) greetMe() {
	fmt.Println(a.name.name, a.name.lastname)
}

//Interfaces

type animal interface {
	makeSound() string
	doAnmialSpecificAction() string
}

type horse struct {
	sound  string
	action string
}

type bird struct {
	sound  string
	action string
}

type fish struct {
	sound  string
	action string
}

//Implemnting interfaces

func (h horse) makeSound() string {
	return h.sound
}

func (h horse) doAnmialSpecificAction() string {
	return h.action
}

func (b bird) makeSound() string {
	return b.sound
}

func (b bird) doAnmialSpecificAction() string {
	return b.action
}

func (f fish) doAnmialSpecificAction() string {
	return f.action
}

func (f fish) makeSound() string {
	return f.sound
}

// Empty interface is used to pass any variable number of paramter.

func variableParam(variableInterface ...interface{}) {
	// underscore (_) is ignoring index value of the array.
	for _, param := range variableInterface {
		paramRecived := param
		fmt.Print(paramRecived)
	}
}

//Pointer Reciver.: you can pass address of the reciver to func if you want to manipulate
// Actual reciever. If address in not passed one copy is created and manipulation is
//Done at that copy.

func (h *horse) pointerReciver() {
	h.sound = "I am chaning horse sound"
	h.action = "i am chaning horse action"
}

func (h *horse) doSomethingElse() {
	h.sound = "hhhh"
	h.action = "action"
}

func main() {
	u := user{"shishir", "dwivedi"}
	x := user{name: "shishir"}
	y := user{lastname: "dwivedi"}
	z := user{}
	u.greetMe()
	x.greetMe()
	y.greetMe()
	z.greetMe()
	p := Person{u}
	p.greetMe()
	h := horse{"Horse Running", "Running"}
	b := bird{"Bird Sound", "flying"}
	f := fish{"fishSound", "swimining"}

	fmt.Println(h.makeSound())
	fmt.Println(h.doAnmialSpecificAction())
	fmt.Println(b.makeSound())
	fmt.Println(b.doAnmialSpecificAction())
	fmt.Println(f.makeSound())
	fmt.Println(f.doAnmialSpecificAction())
	variableParam("shishir", 2, 3+4i, "hello", 3.445665, h, b, f)
	horse := &horse{"Horse", "Run"}
	horse.pointerReciver()
	fmt.Println(*horse)
	horse.doSomethingElse()
	fmt.Println(*horse)
}
