package lesson9

import (
	"fmt"
)

type Circle struct {
	r int
	d int
	c int
}

func Entry() {
	fmt.Println("Structs")

	// var circle1 Circle
	// circle1 := new (Circle)
	circle1 := Circle{
		r: 1,
		d: 1,
		c: 1,
	}

	fmt.Println("circle1", circle1)
	circle1.r = 2
	fmt.Println("circle1", circle1)
}
