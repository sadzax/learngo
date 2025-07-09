package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10)+rand.Intn(10))
	fmt.Printf("My favorite float number is: %.2f", rand.Float64()*100)
}
