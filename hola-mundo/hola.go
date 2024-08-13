package main

import (
	"fmt"
	"math"
	"strconv"

	"rsc.io/quote"
)

const Pi float32 = 3.1415

const (
    X =100
    Y=0b1010//binario
    Z=0x12c//hexadecimal
    A=0o123//octal
)

const (
    Domingo = iota + 1
    Lunes
    Martes
    Miercoles
    Jueves
    Viernes
    Sabado
)
func main() {
    fmt.Println(quote.Hello())
    //declaracion de variables

    var firstName, lastName string //declaracion ed variable del mismo tipo
    var age int

    firstName = "Andres"
    lastName = "Garcia"
    age = 20

    var (
        sex string = "M"
        heigth int = 172
        weigth  = 80// variable sin tipo
    )

    //var isStudent, isWorker, school  = true, false, "UTP" //variables en una sola linea
     isWorker, school  :=  false, "UTP" //variables en una sola linea inicializada

     isStudent := false; //
    

    fmt.Println(firstName, lastName, age,school , isWorker, isStudent, sex, heigth, weigth)

    //declaracion de constantes

    fmt.Println(Pi)
    fmt.Println(X, Y ,Z, A)
    fmt.Println(Viernes)

    //tipos de datos

    var a int8 = 127
    var b float32 = 3.14
    var c uint8 = 255 //solo numeros positivos

    fullName := "Alex Roel \t(alias \"rolecode\")\n)"

    fmt.Println(a, b, c)
    fmt.Println(math.MaxInt) // varl omaxico d3 un int

    fmt.Println(fullName)

    var mybyte byte ='a'
    fmt.Println(mybyte)

    s:="hola"
    fmt.Println(s[0])

    var r rune = 'ⓢ' //alias para int32
    fmt.Println(r);

    //valores predeterminados

    var i int
    var f float64
    var bo bool
    var s2 string
    fmt.Printf("%v %v %v %q\n", i, f, bo, s2)

    //conversiones de tipo

    var integer16 int16 = 50
    var integer32 int32 = 100

    fmt.Println(integer16 + int16(integer32))

    var myString = "100"
    myStringConverted, _ := strconv.Atoi(myString) // con el guion bajo digo que no quiero recibir el error que devuelve la funcion atoi
    fmt.Println(myStringConverted+myStringConverted)

    n:=42;
    s = strconv.Itoa(n) //convierte un entero a string
    fmt.Println(s+s)//concatena dos strings


    var name, apellido string 
    var edad  int
    fmt.Print("ingrese su nombre y epallido:")
    fmt.Scanln(&name, &apellido) //lee desde la consola
    fmt.Print("Ingrese su edad:")
    fmt.Scan(&edad)

    //el paquete fmt
    fmt.Print("otro mensaje") // no hace salto de linea
    fmt.Println("mensaje") //hace salto de linea
    fmt.Printf("tengo  %d y me llamo %s %s \n", edad, name, apellido) //imprime un mensaje formateado

    greting := fmt.Sprintf("hola %s", firstName) //devuelve un string formateado
   
    fmt.Sprintf("hola %v", firstName) //%v cunado no sabemos que tipo de datyo se va a mostrar
  
    fmt.Println(greting)

    fmt.Printf("El tipo de datop de name es: %T\n", name) // con %T averiguo que tipo de dato e suna variable



    //paquete math
    fmt.Println(math.E)
}