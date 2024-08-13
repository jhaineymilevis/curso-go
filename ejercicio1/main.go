package main

import (
	"fmt"
	"math"
)

func main() {
	var side1, side2 float64 
	const presicion =2

	fmt.Print("Ingrese lado 1")
    fmt.Scanln(&side1) 
	fmt.Print("Ingrese lado 2")
    fmt.Scanln(&side2) 

	//procesos
	area := (side1 * side2) /2
	hipotenusa := math.Sqrt(math.Pow(side1, 2) + math.Pow(side2, 2))
	perimetro :=side1+side2+hipotenusa

	fmt.Printf("\n area %.*f: ",presicion,  area)
	fmt.Printf("\n perimetro %.*f: ", presicion, perimetro)


}