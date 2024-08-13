package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//if else

	//t := time.Now()
	//hour := t.Hour()
	//fmt.Println(t)

	//if hour < 12 { //uso de if normal
	if t := time.Now(); t.Hour() < 12 { //uso de if declarando la variable enla condicion, esta variable solo  se puede usar dentro del if
		fmt.Println("es de manana")
	} else if t.Hour() < 17 {
		fmt.Println("Es tarde")
	} else {
		fmt.Println("Es noche")
	}

	//switch

	///os := runtime.GOOS

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println(" Go run -> Windows") //en go no es necesario usar break
	case "linux":
		fmt.Println(" Go run -> Linux")
	case "darwin":
		fmt.Println(" Go run -> macOs")
	default:
		fmt.Println(" Go run -> otro Os")
	}

	//bucle for //es la unica estructura de control repetitiva que existe en go

	var i int
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for j := 1; j <= 10; j++ {

		if j == 3 {
			continue //pasa a la siguiente iteracion
		}

		if j == 5 {
			break //saca del bucle antes de que finalice por completo
		}
		fmt.Println(j)
	}

}
