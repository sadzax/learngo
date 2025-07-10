package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13} // array

	var s []int = primes[0:4] // обозначать []int необязательно
	//var s = primes[1:4] // можно и так
	fmt.Println(s)

	var coolSlice = getArray(primes, 3, 5)
	fmt.Println(coolSlice)
}

// func getArray(arr []int, from int, till int) []int { // Вот так не сработает
func getArray(arr [6]int, from int, till int) []int { // ??? а вот тут уже обязательно обозначит размер
	return arr[from:till]
}
