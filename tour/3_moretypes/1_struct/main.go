package main

import (
	"fmt"
	"reflect"
)

type Human struct { // Struct - это своего рода микрокласс?
	name   string
	age    int
	wealth float64
}

// Как struct хранится в памяти?

func main() {
	guy := Human{"oleg", 2, 300.00}
	guy.age = 4
	fmt.Println(guy.age + 1) // should be 5
	fmt.Println(reflect.TypeOf(guy))
	pointer_to_str(guy)
}

// typeof = main.Human
// здесь main это из package или метода?

func pointer_to_str(h Human) {
	pointer := &h
	pointer.name = "gleb" // почему иногда :=, иногда просто =
	fmt.Println(*pointer)

	fmt.Printf("\n\n")
	fmt.Printf("Type of h: %T\n", h)
	fmt.Printf("Type of pointer: %T\n", pointer)
	fmt.Printf("Type of &h: %T\n", &h)

	fmt.Printf("\n")
	fmt.Println(v1, p, v2, v3)
}

type BodyParams struct {
	height, weight int
}

var (
	v1 = BodyParams{177, 88}    // has type Vertex
	p  = &BodyParams{1, 2}      // Куда ведёт поинтер??!
	v2 = BodyParams{weight: 95} // Y:0 is implicit
	v3 = BodyParams{}           // X:0 and Y:0
)
