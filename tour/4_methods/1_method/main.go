//Go does not have classes. However, you can define methods on types.
//A method is a function with a special receiver argument.
//The receiver appears in its own argument list between the func keyword and the method name.
//In this example, the Abs method has a receiver of type Gauge named g.

package main

import (
	"fmt"
	"math"
	"strconv"
)

type Gauge struct {
	Length, Width float64
}

// Метод Abs для типа Gauge. ///// !!!! ???
// g — это приёмник (receiver), экземпляр Gauge.
// Метод = функция с ресивером!
func (g Gauge) Abs() float64 { // ? надо поставить пробел между func и (), чтобы был именно приёмник?
	return math.Sqrt(g.Length*g.Length + g.Width*g.Width)
}

func Abs2(g Gauge) float64 { // ? надо поставить пробел между func и (), чтобы был именно приёмник?
	return math.Sqrt(g.Length*g.Length + g.Width*g.Width)
}

func (g Gauge) Glow(floatNumber float64) string {
	fmt.Println("in_func_print__ as float32 with 'E' (decimal exponent) :", strconv.FormatFloat(floatNumber, 'E', -1, 32))
	fmt.Println("in_func_print__ as float64 with 'E' (decimal exponent) :", strconv.FormatFloat(floatNumber, 'E', -1, 64))
	fmt.Println("in_func_print__ as float32 with 'f' (no exponent) :", strconv.FormatFloat(floatNumber, 'f', -1, 32))
	fmt.Println("in_func_print__ as float64 with 'f' (no exponent) :", strconv.FormatFloat(floatNumber, 'f', -1, 64))
	fmt.Println("in_func_print__ with fmt.Sprint :", fmt.Sprint(floatNumber))
	fmt.Println("in_func_print__ with fmt.Sprintf :", fmt.Sprintf("%f", floatNumber))
	return fmt.Sprint(floatNumber) // перевод float в string
}

func (g Gauge) Rounder(x int) float64 {
	var floatNumber = math.Abs(g.Length + g.Width + 0.123456789)
	var multiplyer = math.Pow(10.0, math.Abs(float64(x)))
	return math.Round(floatNumber*multiplyer) / multiplyer
}

func (g *Gauge) Scale(x float64) { // зачем нужен поинтер тут?
	g.Length *= x
	g.Width *= x
}

func main() {
	someG := Gauge{3, 4}
	fmt.Println(someG.Abs()) // ==
	fmt.Println(Abs2(someG)) // ==

	fmt.Println(someG.Glow(someG.Rounder(4)))
	fmt.Printf("the type is: %T\n", someG.Glow(someG.Rounder(4)))

	parsed, err := strconv.ParseFloat(someG.Glow(someG.Rounder(4)), 64)
	if err != nil {
		fmt.Println("Ошибка при преобразовании строки в float:", err)
	} else {
		fmt.Printf("parsed float: %f\n", parsed)
		fmt.Printf("type is: %T\n", parsed)
	}
}

// Запись функий, виды:

// 1.1 (через ресивер)
func (g Gauge) SquareRoot() float64 { return math.Sqrt(g.Length*g.Length + g.Width*g.Width) }

// g := Gauge{3, 4}
// fmt.Println(g.SquareRoot())

// 1.2 (классическая запись)
func SquareRoot2(g Gauge) float64 { return math.Sqrt(g.Length*g.Length + g.Width*g.Width) }

// g := Gauge{3, 4}
// fmt.Println(SquareRoot(g))

// 2
func needInt(x int) int { return x*10 + 1 }

// 3
func compute(fn func(float64, float64) float64) float64 { return fn(3, 4) }
