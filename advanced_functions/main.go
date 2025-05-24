package main

import "fmt"

/********** Funciones variadicas: Son funciones que reciben N argumentos  **************/

func sum(name string, nums ...int) int {
	fmt.Printf("%T - %v \n", name, name)
	fmt.Printf("%T - %v \n", nums, nums)
	var total int

	for _, num := range nums {
		total += num
	}
	return total

}

// With this i can receive N data of N type
func printData(data ...interface{}) {

	for _, d := range data {
		fmt.Printf("%T - %v \n", d, d)
	}
}

// recursive functions
func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

// fuction that receive a anonimous funciton and invoke it
func greet(name string, f func(string)) {
	f(name)
}

func duplicate(n int) int {

	return n * 2
}

func triple(n int) int {
	return n * 3
}
func apply(f func(int) int, n int) int {

	return f(n)
}

//function dof superior order, funcione spersonalizadas en tiempo de ejecucion

func double(f func(int) int, n int) int {
	return f(n * 2)
}
func addOne(x int) int {
	return x + 1
}

// clousre
func incrementer() func() int {
	i := 0
	//return a anonimous function
	return func() int {
		//this function can remember the value of i, that was defined in the parent funcion
		i++
		return i
	}
}

func main() {
	sum := sum("Alex", 4, 56, 78, 90, 445)
	fmt.Printf("sum = %d\n", sum)

	printData("HOla", 56, 0, 2, false)

	println(factorial(4))

	// anonimous fucntion
	func() {
		fmt.Println("Hola desde una función anónima")
	}()

	greting := func(name string) {
		fmt.Printf("Hola  %s desde una función anónima\n", name)
	}

	greting("ALex")

	greet("Jhay", greting)

	r1 := apply(duplicate, 5)
	r2 := apply(triple, 5)

	fmt.Println(r1)
	fmt.Println(r2)
	// end anonimous funciton ///

	//functiuons superior order
	r := double(addOne, 3)
	fmt.Println("Result: ", r)

	//closuers, es uan funciona anonima que se definee dentro de otra funcion
	nextInt := incrementer()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}
