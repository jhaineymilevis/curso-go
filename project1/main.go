package main

import (
	"fmt"
	"math/rand"
)

func main() {

	jugar()
}

func jugar() {
	numAleatorio := rand.Intn(100)
	var numeroIngresado int
	var intentos int
	const maxIntentos = 10

	for intentos < maxIntentos {
		intentos++
		fmt.Printf("Ingresa un numero (Intentos restantes : %d):", maxIntentos-intentos+1)
		fmt.Scanln(&numeroIngresado)
		if numeroIngresado == numAleatorio {
			fmt.Println("Felicidade! adivinaste el numero")
			return
		} else if numeroIngresado < numAleatorio {
			fmt.Println("El numero aletaorio es mayor")
		} else if numeroIngresado > numAleatorio {
			fmt.Println("El numero aleatorio es menors")
		}
	}

	fmt.Println("Se acabaron los intentos, el numero era: ", numAleatorio)
	jugarNuevamente()

}

func jugarNuevamente() {
	var eleccion string
	fmt.Println("Quieres Jugar nuevamente? (s/n)")
	fmt.Scan(&eleccion)

	switch eleccion {
	case "s":
		jugar()
	case "n":
		fmt.Println("GRa cias por jugar!!")
	default:
		fmt.Println("eleccion invalida, intentsalo nuevamente")
		jugarNuevamente()

	}
}
