package main

import (
	"errors"
)

func divide(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No se puede dividir por cero")
	}
	return dividendo / divisor, nil

}

/*
func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("resultados:", result)
}*/
