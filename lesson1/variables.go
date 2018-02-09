package lesson1

import (
	"fmt"
	"reflect"
)

// Without initialization variables have 0 (GLOBAL VARIABLES)
var (
	name, course string
	module       float64
	moduleB      = 3.2
)

// Varentry Entry point to the varibales package
func Varentry() {
	fmt.Println("VARIABLES:")

	fmt.Println("Without initialization:", name, course, module)
	fmt.Println("Types:", reflect.TypeOf(name), reflect.TypeOf(course), reflect.TypeOf(module), reflect.TypeOf(moduleB))

	// Invalid operation:
	// a := 10.0000
	// b := 3
	// c := a + b

	// Shorthand declaration
	a := 10.0000
	b := 3
	c := int(a) + b
	// b := c + b Variables at the function level must be used
	fmt.Println("c=", c)

	// POINTERS:
	ptr := &c
	fmt.Println("Memory address of *c* is", ptr)
	fmt.Println("*c* by ptr:", *ptr)
}
