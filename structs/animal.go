package structs

type Animal struct {
	name string
	age  int
}

func NewAnimal(name string, age int) Animal {
	newAnimal := Animal{name, age}
	return newAnimal
}

func (varAnimal Animal) Print() string {
	return varAnimal.name
}
