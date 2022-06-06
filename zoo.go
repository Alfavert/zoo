package main

import "fmt"

type Animal interface {
	MakeSound()
	Eat() string
}

type Dog struct {
	Sound string
	Food  string
}
type Fish struct {
	Sound string
	Food  string
}
type Cat struct {
	Sound string
	Food  string
}

func (dog Dog) MakeSound() {
	fmt.Println(dog.Sound)
}
func (dog Dog) Eat() string {
	return fmt.Sprintf("%s ", dog.Food)
}
func (fish Fish) MakeSound() {
	fmt.Println(fish.Sound)
}
func (fish Fish) Eat() string {
	return fmt.Sprintf("%s ", fish.Food)
}
func (cat Cat) MakeSound() {
	fmt.Println(cat.Sound)
}
func (cat Cat) Eat() string {
	return fmt.Sprintf("%s ", cat.Food)
}
func main() {
	slice := []Animal{Dog{"Woof", "Bones"}, Fish{"0.o", "worm"}, Cat{"Meow", "Fish"}}
	var food string
	for i := range slice {
		slice[i].MakeSound()
		food += slice[i].Eat()
	}
	fmt.Println(food)
}
