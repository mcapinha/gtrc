package models

import "fmt"

func SayHello() {
	fmt.Println("Hello everyone from package1")
}

type Tuple struct {
	x, y, z float64
	w       int8
}

func Point(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1}
}

func Vector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0}
}
