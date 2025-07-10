package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13} // array

	var s []int = primes[0:4] // обозначать []int необязательно
	//var s = primes[1:4] // можно и так
	fmt.Println(s)

	var coolSlice = getArray(primes, 3, 5)
	fmt.Println(coolSlice)

	fmt.Println(len(coolSlice))
	fmt.Println(cap(coolSlice))

	beatlesNames()
}

// func getArray(arr []int, from int, till int) []int { // Вот так не сработает
func getArray(arr [6]int, from int, till int) []int { // ??? а вот тут уже обязательно обозначит размер
	return arr[from:till]
}

func beatlesNames() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	// Почему вот тут будет размерность?
	fmt.Printf("And it's a: %T\n", names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println("\n", a, b)
	// А вот тут её не будет, ведь она определена
	fmt.Printf("And it's a: %T, b: %T\n", a, b) // "And it's a: []string, b: []string"

	b[0] = "XXX"
	a[1] = "YYY" // переопределит в обоих местах
	fmt.Println(a, b)
	fmt.Println(names)
}

// Slice - для чего в целом? Это как "пойнтер" в своём роде, но для Array? Slice не создаёт новый набор данных, а оперирует с уже существующим?
