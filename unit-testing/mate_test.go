package main

import "testing"

func TestSumar(t *testing.T) {

	/*result := Sumar(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Error al sumar: se esperaba %d, pero se obtuvo %d", expected, result)
	}*/

	casos := []struct {
		a, b     int
		expected int
	}{
		{2, 3, 5},
		{0, 0, 0},
		{-1, 1, 0},
		{-2, -3, -5},
		{100, 200, 300},
	}

	for _, caso := range casos {
		result := Sumar(caso.a, caso.b)
		if result != caso.expected {
			t.Errorf("Error al sumar %d y %d: se esperaba %d, pero se obtuvo %d", caso.a, caso.b, caso.expected, result)
		}
	}

}
