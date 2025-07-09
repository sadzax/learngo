package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z float64 = 0
	for math.Pow(z, 2.0) < x {
		powered := math.Pow(z, 2.0)
		fmt.Printf("The powered %v in 2 is %v\n", z, powered)
		z++
	}

	prev := math.Pow((z - 1), 2.0)
	prev_bend := math.Abs(prev - x)
	next := math.Pow(z, 2.0)
	next_bend := math.Abs(next - x)

	fmt.Printf("The answer prev is %v\n", (z - 1))
	fmt.Printf("The answer next is %v\n", z)

	if next_bend < prev_bend {
		return z
	}

	return z - 1
}

func main() {
	fmt.Println(Sqrt(34))
}
