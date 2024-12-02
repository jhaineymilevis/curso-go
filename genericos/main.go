package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func PrintList(list ...any) {
	for _, item := range list {
		fmt.Println(item)
	}

}

type integer int

type Numbers interface {
	~int | ~float32 | ~float64 | ~uint
} //crear restricccion

// funcion generica, con parametro de tipo, restricicones de apromaximacion
// func Sum[T ~int | ~float64](nums ...T) T {
// func Sum[T Numbers](nums ...T) T {
func Sum[T constraints.Integer | constraints.Float](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}

	return total
}

func main() {
	PrintList("Alex", "John", "Jhainey")
	PrintList(3, 6, 7, 8, 9, 10, 11)
	PrintList("hola", 6, 7, 8, 9, 10, 11)

	fmt.Println(Sum[int](4, 3, 4, 5, 3))
	fmt.Println(Sum[float64](4.5, 3.1, 4.4, 5, 3.6))

	var myInt integer = 5
	var myInt2 integer = 300
	fmt.Println(myInt)
	fmt.Println(myInt2)
	fmt.Println(Sum(myInt, myInt2))
}
