// MAP - это хеш-таблица?

package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

type Person struct {
	name   string
	age    uint8
	weight float64
}

func main() {
	createAndPrintHRManager()
	fmt.Println(musicBand)
	fmt.Printf("%T\n\n", musicBand)
	deleteElemFromMap()
	fmt.Printf("\n\n")
	exersize()
}

func createAndPrintHRManager() {
	m := make(map[string]Person)
	m["Our HR Manager"] = Person{
		"Olga",
		23,
		50.23,
	}
	fmt.Println(m["Our HR Manager"])
}

var musicBand = map[string]Person{
	"Guitar": {
		"Vova", 45, 85.3,
	},
	"Drums": {
		"Kirill", 23, 75.2,
	},
}

func deleteElemFromMap() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	fmt.Println("The map:", m)

	v, ok := m["SOME_KEY"] // Второй аргумент примет наличие/отсутствие нахождения ключа
	fmt.Println("The value:", v, "Present?", ok)
}

func exersize() {
	wc.Test(WordCount)
}

func WordCount(s string) map[string]int {
	res := make(map[string]int)

	// Вводить сюда переменную -- это не Go-way? Можно просто strings.Fields(s)
	arrayed_s := strings.Fields(s)

	fmt.Println(arrayed_s)

	for i := range arrayed_s { // только итератор
		if res[arrayed_s[i]] == 0 {
			res[arrayed_s[i]] = 1
		} else {
			res[arrayed_s[i]] = +res[arrayed_s[i]]
		} // Вот тут не складывается, как будто `res` нужно возвращать

		//res[arrayed_s[string(i)]] = 1 + int(res[arrayed_s[i]]) // почему тут не работает?

		//res[arrayed_s[i]] = 1 + int(res[arrayed_s[i]])
	}

	return res
}
