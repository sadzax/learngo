package main

import (
	"fmt"
	"math"
)

// compute принимает функцию "fn" как аргумент
// fn должна принимать два float64 и возвращать 1 float64
// compute тоже вернёт 1 float64
func compute(fn func(float64, float64) float64) float64 {
	// Вызываем переданную функцию с аргументами 3 и 4
	return fn(3, 4)
}

func main() {
	// Объявляем функцию  и присваиваем её переменной  hypot
	// Это анонимная функция (lambda) для нахождения гипотенузы
	hypot := func(x, y float64) float64 {
		karateSum := float64(math.Pow(x, 2) + math.Pow(y, 2))
		return math.Sqrt(karateSum)
	}

	// Вызываем hypot напрямую с аргументами 5 и 12.
	fmt.Println(hypot(5, 12)) // Выведет 13

	// Передаём hypot как аргумент в compute.
	fmt.Println(compute(hypot)) // Выведет 5

	// Передаём стандартную функцию math.Pow как аргумент в compute.
	fmt.Println(compute(math.Pow)) // Выведет 81 (3^4)

	funcClosure() // + ещё Замыкание
}

// Замыкание — это функция, которая «запоминает» переменные из внешнего окружения
// , даже если вызывается вне этого окружения.
// FUNC_CLOSURE
// adder возвращает функцию (замыкание), которая увеличивает сумму на x.
func adder() func(int) int {
	// ?? = ПОЧЕМУ НЕ СБРАСЫВАЕТСЯ НА 0, ведь ADDER в себе содержит это
	sum := 0 // переменная из внешнего окружения для замыкания
	return func(x int) int {
		sum += x // используем и изменяем внешнюю переменную sum
		return sum
	}
}

func funcClosure() {
	pos, neg := adder(), adder() // два независимых замыкания ?
	// У каждого своя собственная переменная sum.

	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),    // увеличивает свою сумму на i
			neg(-2*i), // увеличивает свою сумму на -2*i
		)
	}
}
