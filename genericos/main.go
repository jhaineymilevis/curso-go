package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// asi se hacia antes
func PrintListOldWay(list ...interface{}) {
	for _, item := range list {
		fmt.Println(item)
	}

}

// despues de go 1.18 se usa any
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

func Includes[T comparable](list []T, value T) bool {

	for _, item := range list {
		if item == value {
			return true
		}
	}

	return false
}

func Filter[T constraints.Ordered](list []T, callback func(T) bool) []T {
	filtered := make([]T, 0, len(list))

	for _, item := range list {
		if callback(item) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

func main() {
	PrintList("Alex", "John", "Jhainey")
	PrintList(3, 6, 7, 8, 9, 10, 11)
	PrintList("hola", 6, 7, 8, 9, 10, 11)

	fmt.Println(Sum(4.7, 3.6, 4, 5, 3))
	fmt.Println(Sum[int](4, 3, 4, 5, 3))
	fmt.Println(Sum[float64](4.5, 3.1, 4.4, 5, 3.6))

	var myInt integer = 5
	var myInt2 integer = 300
	fmt.Println(myInt)
	fmt.Println(myInt2)
	fmt.Println(Sum(myInt, myInt2))

	strings := []string{"a", "b", "c"}
	numbers := []int{1, 2, 3, 6, 7}

	fmt.Println(Includes(strings, "f"))
	fmt.Println(Includes(numbers, 3))

	fmt.Println(Filter(numbers, func(value int) bool { return value > 3 }))
	fmt.Println(Filter(strings, func(value string) bool { return value > "b" }))
}
