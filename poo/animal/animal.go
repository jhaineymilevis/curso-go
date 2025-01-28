package animal

import "fmt"

type Animal interface {
	Sound()
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

// implement Animal inteface
func (d *Dog) Sound() {
	fmt.Println("Woof!")
}

// implement Animal inteface
func (d *Cat) Sound() {

	fmt.Println("Mew!")
}

/** This function decide with sound makes depending of the animal type **/
func MakeSound(animal Animal) {

	animal.Sound()

}
