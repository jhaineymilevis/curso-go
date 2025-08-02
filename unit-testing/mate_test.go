package main

import "testing"

func TestSumar(t *testing.T) {

	result := Sumar(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Error al sumar: se esperaba %d, pero se obtuvo %d", expected, result)
	}
}
