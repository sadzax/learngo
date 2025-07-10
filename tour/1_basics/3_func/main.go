package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

//  Почему вот так нельзя? https://go.dev/tour/basics/7
//  func swap(x, y string) (y, x string) {
//  	return
//  }

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
