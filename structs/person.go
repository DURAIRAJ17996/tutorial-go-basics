package structs

type person struct {
	name   string
	age    int
	isGood bool
}

func newPerson(name string, age int, isGood bool) *person {
	varPerson := person{name, age, isGood}
	return &varPerson
}
