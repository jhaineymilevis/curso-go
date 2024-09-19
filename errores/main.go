package main

import (
	"fmt"
	"log"
	"os"
)

func divide(dividendo, divisor int) (int, error) {
	//funcion anonima
	defer func() { // con defer se ejecuta al final de la cuncion divide
		//implementamos recover, con esto se logra que el panic no interrumpa nuestra aplicacion
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	//con trol de erroes
	/*if divisor == 0 {
		return 0, errors.New("No se puede dividir por cero")
	}*/
	validateZero(divisor) //control de erroes con panic
	return dividendo / divisor, nil

}

func validateZero(divisor int) {

	if divisor == 0 {
		panic("No se puede dividir por cero") //panic
	}
}

func main() {

	//errores
	result, err := divide(10, 2)
	// paninc y recover . no se recomienda para el flujo normal de control del programa. se recomienda para casos excepcionales o errores graves o situaciones inesperdas. manejar limpieza antyes de finalizar el programa
	divide(10, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("resultados:", result)

	//defer
	defer fmt.Println(3) //esta funcion se ejecuta al final
	fmt.Println(2)
	fmt.Println(1)

	//crear archivo
	file, err := os.Create("hola.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	//cerrar el archivo con defer para que aunque falle lo demas, siempre se ejecute al final
	defer file.Close()

	//escribir en el archivo
	_, err = file.Write([]byte("hola mundo"))
	if err != nil {
		fmt.Println(err)
		return
	}

	//regstro de erores
	log.SetPrefix("main(): ") // agrega un prefijo
	log.Println("ejecutando progrma")

	//regiustrar los erroes en un archiuvo
	file2, err2 := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err2 != nil {
		log.Fatal(err2)
	}
	defer file2.Close()
	log.SetOutput(file2)
	log.Print("oye, soy un log")

	//log.Fatal("error en la aplicacion fatal, esto detiene el programa")
	//log.Panic("error en la aplicacion panico, esto detiene el programa")
}
