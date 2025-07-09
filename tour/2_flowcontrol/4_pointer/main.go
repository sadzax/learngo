package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j       // point to j
	*p = *p / 37 // divide j through the pointer

	fmt.Println(j) // see the new value of j

	fmt.Println(i) // Пойнтер отвязывается от старой ячейки памяти?

	play_with_pointer()
}

func play_with_pointer() {
	kek := "cheburek"
	//kek := 'cheburek' // сингл-кавычки не используем?
	lol := &kek
	fmt.Println(*lol)
}
