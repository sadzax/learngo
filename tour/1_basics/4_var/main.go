package main

//import "fmt"
//import "reflect"
// Можно и так

import (
	"fmt"
	"reflect" // reflect.TypeOf выведет тип объекта/переменной
)

func main() {
	var i, j float64 = 1, 2 // set floats
	k := 3                  // set auto (int)
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(j))
	fmt.Println(reflect.TypeOf(k))
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(python))
	fmt.Println(reflect.TypeOf(java))
}
