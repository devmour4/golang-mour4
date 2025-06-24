package main

import (
	"fmt"
)
//utilize seus conhecimentos de Go até agora para implementar uma calculadora 
//completa: a função de soma, subtração, adição e multiplicação e apresente os resultados corretos na tela.
func main() {
	w := Sum(1, 2, 3)
	x := Sub(5, 10)
	y := Mult(10, 10)
	z := Division(20)

	fmt.Println(w, x, y, z)
	
}

func Sum (i ...int) int {
	total := 0
	
	for _, v := range i {
		total += v
	}
	return total
}
func Sub (i	...int) int {
	total := 0

	for _, v := range i {
		total = v - total
	}
	return total
}
func Mult (i ...int) int {
	total := 1
	
	for _, v := range i {
		total = total * v
	}
	return total
}
func Division (i ...int) int {
	total := 1
	
	for _, v := range i {
		total = v / total
	}
	return total
}