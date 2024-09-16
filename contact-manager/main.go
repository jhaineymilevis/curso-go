package main

import (
	"fmt"
	"os"
)

func main() {
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

}
