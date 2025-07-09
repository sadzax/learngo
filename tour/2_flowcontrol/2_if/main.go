package main

import (
	"fmt"
	"math"
	"time"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	//  Это не два условия, а инициализация + условие.
	//  if <инициализация>; <условие>
	//  `v` здесь - НЕ ДОСТУПНО
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		string("\n"),
		pow(3, 3, 20),
		string("\n"),
		string("\n"),
	)

	// + switch
	check_time()

	// + defer
	check_backwards()
}

func check_time() {
	// + switch
	fmt.Printf("Right now it's %v \n\n", time.Now())
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	fmt.Println(time.Now().Weekday())
}

func check_backwards() {
	// + defer
	the_time := time.Now()
	fmt.Printf("\n\n\n Ah! By the way tomorrow will be %v \n\n", the_time.Weekday()+1)

	fmt.Println("\n And countdown is a:")
	for i := 0; i < 10; i++ {
		k := 10 - i
		// fmt.Printf("\n %v", the_time.Weekday()+k) // Почему тут не работает?
		fmt.Println(k)
		defer fmt.Println(k) // не обязательно в "функции" - но и в "цикле" / "проке"
	}

	fmt.Println("\n And the backwards again:")
	// В Go ключевое слово defer используется для отложенного выполнения функции. Это значит, что функция, объявленная с помощью defer, будет вызвана только после завершения окружающей функции (той, в которой находится defer).
}
