package main

import (
	"fmt"
	"math"
)

// var powArray = []int{1, 2, 4, 8, 16, 32, 64, 128}
var powArray = calculateThePowArray(7)

func main() {
	// The range form of the for loop iterates over a slice or map.
	iterateLikeARange()
}

func iterateLikeARange() {
	for i, v := range powArray { // здесь каждый элемент поочерёдно попадает в итератор v ?
		// Это не то же, что и циклы, нет ";"
		fmt.Printf("2**%d = %d \n", i, v)
	}
}

func calculateThePowArray(numb int) []float64 {
	var sum []float64
	for i := 0; i < numb; i++ {
		res := math.Pow(2, float64(i))
		sum = append(sum, res)
	}
	return sum
}

func printKek() {
	fmt.Printf("haha")
}
