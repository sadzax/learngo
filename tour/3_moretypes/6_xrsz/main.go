package main

import "golang.org/x/tour/pic"

// uint8 - практически для чего используется? условные IP
// Почему CAMEL-CASE? Эту функция будет экспортироваться?
func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy) // матрица
	for y := 0; y < dy; y++ {
		row := make([]uint8, dx) // внутренняя строка
		for x := 0; x < dx; x++ {
			row[x] = uint8(x ^ y)
			// побитовая операцию XOR (исключающее ИЛИ) ?? для чего это в целом?
		}
		pic[y] = row
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
