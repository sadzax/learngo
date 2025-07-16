package main

import "fmt"

func main() {
	createNilArray()

	makeSlice()
}

func createNilArray() {
	var megaList []int
	// fmt.Println(megaList, len(megaList), cap(megaList))
	if megaList == nil { // не объявленные переменные != nil
		// невозможно слайсить nil-array
		fmt.Println("Created array is really nil!")
	} else {
		fmt.Println("Created array is not nil at all") // сюда не дойдём
	}
}

// Чем в Go являетяс nil? Это объект / класс / тип?

// Creating a slice with make
// Называется  Creating a slice with make - а создаются не слайсы, а Array, разве нет?
func makeSlice() {
	a := make([]int, 5) // len(b)=0, cap(b)=5
	printSlice("a", a)

	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) { // Хочу импортировать этот метод из ./tour/3_moretypes/3_slice/main.go
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
	fmt.Printf("type is: %T\n", x)
}
