package main

import (
	"fmt"
	"reflect"
)

func main() {
	var words [2]string
	words[0] = "Hello" // опять же, тут без :=
	words[1] = "World" // опять же, тут без :=
	fmt.Println(words)
	fmt.Println(words[0], "\n", words[1])
	fmt.Printf("\n%T\n", words)        // [2]string -- array - не пишется
	fmt.Println(reflect.TypeOf(words)) // [2]string -- array - не пишется

	//primes := []int{2, 3, 5, 7, 11, 13} // столько и будет, если не указать длину
	primes := [6]int{2, 3, 5, 7, 11, 13}
	primesLong := [7]int{2, 3, 5, 7, 11, 13} // непереданный элемент = Zero Values
	//primes_short := [5]int{2, 3, 5, 7, 11, 13} // !!! ERR
	fmt.Println(primes, "\n", primesLong)
}
