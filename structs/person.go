package structs

/*
Golang struct is equivalent to DTO in java.
Constructor for struct is advised to have the naming conventions as
New<structName>
*/
type Person struct {
	name   string
	age    int
	isGood bool
}

// Exported function, can be accessed outside this package
func NewPerson(name string, age int, isGood bool) Person {
	varPerson := Person{name, age, isGood}
	return varPerson
}

// Un-exported function, cannot be accessed outside this package
func newPerson(name string, age int, isGood bool) Person {
	varPerson := Person{name, age, isGood}
	return varPerson
}

// Method
func (person Person) Print() string {
	return person.name
}
