package main

import "fmt"

func main() {
	fmt.Println(sumToTheNumber(19))
	fmt.Println(sumToTheNumberCont(1000))
	fmt.Println(sumToTheNumberCont2(10))
}

func sumToTheNumber(counter int) int {
	sum := 0
	for i := 0; i < counter; i++ {
		sum += i
	}
	return sum
}

func sumToTheNumberCont(counter int) int {
	sum := 1
	for sum < counter {
		sum += sum
	}
	return sum
}

func sumToTheNumberCont2(counter int) int {
	sum := 0
	i := 0
	for i < counter {
		sum += i
		i++
	}
	return sum
}

// В Go можно использовать for с одной, двумя, тремя или вообще без выражений.
// Главное — соблюдать синтаксис: точки с запятой нужны, если вы явно пишете init и/или post.
// Скобки вокруг выражений не нужны, а фигурные скобки вокруг тела цикла обязательны.
