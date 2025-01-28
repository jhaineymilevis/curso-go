package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tarea struct {
	nombre     string
	desc       string
	completado bool
}

type ListaTareas struct {
	tareas []Tarea
}

// metodo para agregar tarea
func (lt *ListaTareas) AgregarTarea(t Tarea) {
	lt.tareas = append(lt.tareas, t)
}

// marca una tarea como completado
func (lt *ListaTareas) marcarCompletado(index int) {
	lt.tareas[index].completado = true
}

// funcion para editar una tarea
func (lt *ListaTareas) editarTarea(index int, t Tarea) {
	lt.tareas[index] = t
}

// funcion para editar una tarea
func (lt *ListaTareas) eliminarTarea(index int) {
	lt.tareas = append(lt.tareas[:index], lt.tareas[index+1:]...)
}

func main() {

	lista := ListaTareas{}

	//instancia de bufio para entrada de datos
	leer := bufio.NewReader(os.Stdin)

	for {
		var opcion int
		fmt.Println("Seleccione una opcion: \n",
			"1. Agregar tarea \n",
			"2. Mrcar tarea como completada\n",
			"3. Editar tarea \n",
			"4. Eliminar tarea \n",
			"5. Salir \n",
		)
		fmt.Println("Ingrese un opcion")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			var t Tarea
			fmt.Print("Ingrese el nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n') // el gion bajo quiere decir que no voy a leer el error que retorna la funcion
			fmt.Print("Ingrese la descripcion de la tarea: ")
			t.desc, _ = leer.ReadString('\n')
			lista.AgregarTarea(t)
			fmt.Println("Tarea agregada exitosamente!")

		case 2:
			var index int
			fmt.Print("Ingrese el indice de la tarea a marcar como completada: ")
			fmt.Scanln(&index)
			lista.marcarCompletado(index)

		case 3:
			var index int
			var t Tarea
			fmt.Println("Ingresa el indice de la tarea a editar")
			fmt.Scanln(&index)
			fmt.Print("Ingrese el nuevo nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Print("Ingrese la nueva descripcion de la tarea: ")
			t.desc, _ = leer.ReadString('\n')
			lista.editarTarea(index, t)
			fmt.Println("Tarea editada exitosamente!")
		case 4:
			var index int
			fmt.Print("Ingrese el indice de la tarea a eliminar: ")
			fmt.Scanln(&index)
			lista.eliminarTarea(index)
			fmt.Println("Tarea eliminada exitosamente!")
		case 5:
			fmt.Println("Saliendo del programa...")
			return

		default:
			fmt.Println("Opcion invalida!")

		}

		fmt.Println("Lita de tareas:")
		fmt.Println("================================================================")

		for i, t := range lista.tareas {
			fmt.Printf("%d. %s - %s - Completado %t\n", i, t.nombre, t.desc, t.completado)
		}

		fmt.Println("================================================================")
	}

}
