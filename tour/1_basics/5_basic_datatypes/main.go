package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// В Go `<<` — это только оператор сдвига битов и больше нигде не применяется

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	convert()

	//j := check_new_var_mark(3)
	//fmt.Println(j)
	//fmt.Println(reflect.TypeOf(j))
}

func convert() { // Type conversions
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

//func check_new_var_mark(i, int) int {
//	var i float64 = float64(i)
//	var j float64 = i / 2 // j is an int ?
//	var j int = int(j)
//	return j
//}

// Почему вот так нельзя?
