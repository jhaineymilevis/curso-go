package main

import "fmt"

func main() {
	name := "jhainey"
	saludo := hello(name)
	fmt.Println(saludo)
	sum, mul := calc2(2, 3)
	println("La suma es :", sum)
	println("La multiplicacion  es :", mul)
}

func hello(name string) string {

	return "Hello " + name
}

func calc(a, b int) int { //los dos parametros son del mismo tipo
	return a + b
}

func calc2(a, b int) (int, int) { //puede devolver dos valores
	sum := a + b
	mul := a * b

	return sum, mul
}

func calc3(a, b int) (sum, mul int) { //de estea menraa e ssimilar a la de arriba pero de uan vez declaro las variables que va a retornar
	sum = a + b
	mul = a * b

	return //no hace falta edcirle las variables porque arriba ya sabe que variables va a retornar
}
