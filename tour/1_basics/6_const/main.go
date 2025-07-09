package main

import (
	"fmt"
	"reflect"
)

const Pi = 3.14

// Константы просто пишутся с большой буквы? Зачем?
// Как их отличить в коде от export- данных?

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println("\n\n\n\n")

	check_more_numeric_constants()
}

const ( // запись присвоения объектов через () - что ещё можно так присвоить?
	Big   = 1 << 100
	Small = Big >> 99
)

var (
	superkek = 1.99
	// superkek := 1.99 Почему именно в скобках не работает присвоение через :=   ?
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func check_more_numeric_constants() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	fmt.Printf("\n\n\n\n") // Работает ли Printf без аргументов?
	fmt.Println(superkek)
	fmt.Println(reflect.TypeOf(superkek))
}
