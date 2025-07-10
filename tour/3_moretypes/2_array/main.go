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
	fmt.Printf("\n%T\n", words)        // [2]string -- (array) - не пишется
	fmt.Println(reflect.TypeOf(words)) // [2]string -- (array) - не пишется

	primes := [6]int{2, 3, 5, 7, 11, 13}
	primesUndefinedLengh := []int{2, 3, 5, 7, 11, 13} // столько и будет, если не указать длину
	primesLong := [7]int{2, 3, 5, 7, 11, 13}          // непереданный элемент = Zero Values
	//primes_short := [5]int{2, 3, 5, 7, 11, 13} // !!! ERR

	printArr(primesUndefinedLengh)
	//printArr(primes) 		// не запустится, т.к. передаёт [6]int туда, где надо []int
	//printArr(primesLong)	// не запустится, т.к. передаёт [7]int туда, где надо []int

	fmt.Printf(" \n\n primes: \n")
	fmt.Printf("len=%d cap=%d %v\n", len(primes), cap(primes), primes)
	fmt.Printf("type is: %T\n", primes)

	fmt.Printf(" \n\n primesLong: \n")
	fmt.Printf("len=%d cap=%d %v\n", len(primesLong), cap(primesLong), primesLong)
	fmt.Printf("type is: %T\n", primesLong)

}

func printArr(arr []int) { // тут надо всегда указать точное количество элементов? Иначе как запускать?
	fmt.Printf("\n\n printArr return:\n")
	fmt.Printf("len=%d cap=%d %v\n", len(arr), cap(arr), arr)
	fmt.Printf("type is: %T\n", arr)
}
