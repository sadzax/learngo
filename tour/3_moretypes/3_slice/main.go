// Slice - для чего в целом? Это как "пойнтер" в своём роде, но для Array? Slice не создаёт новый набор данных, а оперирует с уже существующим?

package main

import (
	"fmt"
	"strings"
)

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13} // array

	var s []int = primes[0:4] // обозначать []int необязательно
	//var s = primes[1:4] // можно и так
	fmt.Println(s)

	var coolSlice = getArray(primes, 3, 5)
	fmt.Println(coolSlice)

	fmt.Println(len(coolSlice))
	fmt.Println(cap(coolSlice))

	fmt.Println("\n NEXT BLOCK beatlesNames:")
	beatlesNames()

	fmt.Println("\n NEXT BLOCK getSliceLiteral:")
	getSliceLiteral()

	fmt.Println("\n NEXT BLOCK sliceDefault:")
	sliceDefault()

	fmt.Println("\n NEXT BLOCK sliceLenAndCap:")
	sliceLenAndCap()

	fmt.Println("\n NEXT BLOCK ticTacToe:")
	ticTacToe()

	fmt.Println("\n NEXT BLOCK appendToSlice:")
	appendToSlice()
}

// func getArray(arr []int, from int, till int) []int { // Вот так не сработает
func getArray(arr [6]int, from int, till int) []int { // ??? а вот тут уже обязательно обозначит размер
	return arr[from:till]
}

func getSliceLiteral() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	fmt.Printf("\n Type is: %T\n", q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)
	fmt.Printf("\n Type is: %T\n", r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
	fmt.Printf("\n Type is: %T\n", s) //  Type is: []struct { i int; b bool } - каждый внутренний элемент struct(литерал\ключ) должен быть определён?
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

func sliceDefault() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	new_s := []int{2, 3, 5, 7, 11, 13}

	new_s = new_s[:]
	fmt.Println(new_s)
}

func sliceLenAndCap() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)
	//len=6 cap=6 [2 3 5 7 11 13]
	//type is: []int

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)
	//len=0 cap=6 []
	//type is: []int

	// Extend its length.
	s = s[:4] // несмотря на переопределение s выше в len=0 cap=6 [] Go для слайса всё равно взял самый первый array, но...
	printSlice(s)
	//len=4 cap=6 [2 3 5 7]
	//type is: []int

	// Drop its first two values.
	s = s[2:] // ... тогда почему здесь не [5 7 11 13]? (см. выше пример)
	printSlice(s)
	//len=2 cap=4 [5 7]
	//type is: []int
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Printf("type is: %T\n", s)
}

func ticTacToe() {
	// Slice of slices
	board := [][]string{
		[]string{"_", "_", "_"}, // Почему тут "string" ?
		[]string{"_", "_", "_"}, // как в теории сюда положить другой тип данных?
		[]string{"_", "_", "_"},
	}

	fmt.Println(board)
	fmt.Printf("\n the type is %T\n", board)

	cell00 := &board[0][0]
	*cell00 = "X"

	cell22 := &board[2][2]
	*cell22 = "E"

	// board[2][2] = "K"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func appendToSlice() {
	var s []int // опять же, как отличить array от slice?
	printSlice(s)

	// Результат аппенда - новый объект?
	s = append(s, 0)
	printSlice(s)

	s = append(s, 1) // можно ли выбрать индекс, куда вместить элемент?
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)
}
