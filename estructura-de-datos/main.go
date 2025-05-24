package main

import "fmt"

func main() {
	//matrices()
	//sslices()
	//mapas()
	//estructuras()
	punteros()
}

// estructura de datos fijas
func matrices() {
	var a [5]int                    //matrix: /onjunto de elementos fijo del mismo tipo
	var b = [5]int{1, 2, 3, 4, 5}   //matriz inicializada
	var c = [...]int{1, 2, 3, 4, 5} //matriz inicializada sin saber la cantidsad de elementos
	a[0] = 10
	a[1] = 20
	a[2] = 30
	a[3] = 40
	a[4] = 50
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(a[1])
	for i := 0; i < len(a); i++ {
		fmt.Println(i)
	}
	//otra forma de iterar una matriz
	for index, value := range a {
		fmt.Printf("indice: %d, Valor: %d", index, value)
		//matriz bidimencional
		var matrizBi = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		fmt.Println(matrizBi)

	}

}

// estructura de datos dinamicas
func slices() {
	var a []int // slice sin inicializar
	fmt.Println(len(a))
	diasSemana := []string{"Domingo", "Lunes", "Martes", "Miercoles", "Jueves", "viernes", "sabado"}
	//slice inicializado
	diaRebanada := diasSemana[0:5] //imprime del 0 al 5 elementos del array
	fmt.Println(diasSemana)
	fmt.Println(diaRebanada)
	diaRebanada = append(diaRebanada, "Viernes", "Sabado", "Otro dia") //agrego un elmento al arreglos, si le agrego "otro dia", aumenta la capacidad
	diaRebanada = append(diaRebanada[:2], diasSemana[3:]...)           //elimino el indice 2 del arreglo usando la funcion append
	fmt.Println(len(diaRebanada))                                      //len de la rebnada
	fmt.Println(cap(diaRebanada))                                      //capacidad de almacenamiento de la rebanada
	nombres := make([]string, 5, 10)                                   //creo una rebanada con la funcion make
	nombres[0] = "alex"
	//nombres[1] = "marian" //! esto no se puede, se debeusar la funcion append
	fmt.Println(nombres)

	rebanada1 := []int{1, 2, 3, 4, 5}
	rebanada2 := make([]int, 5)

	copy(rebanada2, rebanada1) //copio los valkores de rebanada1 en rebanada2

	fmt.Println(rebanada1)
	fmt.Println(rebanada2)
}

func mapas() {
	colors := map[string]string{
		"rojo":     "#FF0000",
		"azul":     "#0000FF",
		"amarillo": "#FFFF00",
	}

	fmt.Println(colors["rojo"])

	colors["negro"] = "000000"

	fmt.Println(colors)
	negro, ok := colors["negro"] //el ok es una verificacion, me devuelve un booleano, donde true es que el valro existe en el mapa
	fmt.Println(negro, ok)

	if ok {
		fmt.Println("existe el valoe")
	} else {
		fmt.Println("no existe el valor")
	}

	//lo mismo que arriba pero una forma abreviada de validar y tomar accion en funcion de existe o no la clave
	if verde, okVede := colors["verde"]; okVede {
		fmt.Println(verde)
	} else {
		fmt.Println("no existe el valor")
	}

	//eliminar un elemento
	delete(colors, "negro")
	fmt.Println(colors)

	//iterar sobre un mapa
	for clave, valor := range colors {
		fmt.Printf("Clave: %s, valor: %s\n", clave, valor)
	}

}

// el equivalente a interfaces o types

type Persona struct {
	nombre string
	edad   int
	correo string
}

// es un metodo de persona
func (p *Persona) diHola() {
	fmt.Println("Hola mi nombre es ", p.nombre)
}
func estructuras() {
	var juan Persona
	juan.nombre = "Juan"
	juan.edad = 25
	juan.correo = "juan@example.com"

	fmt.Println(juan)

	maria := Persona{"Maria", 20, "maria@example.com"}

	maria.edad = 25
	fmt.Println(maria)
	fmt.Println(maria.nombre)
}

// almacenann la direccion en memoria de una variable
func punteros() {
	var i int = 10
	var p *int = &i //p es un puntero a la variable i

	fmt.Println("valor de i: ", i)
	fmt.Println("valor de i: ", &i) //referncia de memoria de i
	fmt.Println("valor de p: ", p)

	editar(&i)
	fmt.Println("valor de i: ", i)

	alex := Persona{"Alex", 28, "alex@gmail.com"}
	alex.diHola()
}

func editar(i *int) { //recibe un puntero
	*i = 2 //modifica la posicion en memoria de i

}
