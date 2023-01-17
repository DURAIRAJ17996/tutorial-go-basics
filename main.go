package main

import (
	"fmt"

	"github.com/DURAIRAJ17996/tutorial-go-basics/mock"
)

func main() {
	//fmt.Println(variables.Test())
	// datastructure.TestMapWithRange()
	// Multiple return values
	//index, indexValue := functions.FunctionWithTwoRgo beturn(3)
	//fmt.Println(index, indexValue)
	// varStringList := []string{"India", "Sri Lanka", "Maldives"}
	// functions.FunctionWithVariableArgs(varStringList...)

	/*
		nextSeq := functions.FunctionWithAnonymousFunc()
		fmt.Println(nextSeq())
		fmt.Println(nextSeq())
		fmt.Println(nextSeq())
	*/
	//structs.NewPerson("MethodTEst", 12, true)

	/* Implementing interface
	var interfaceObj structs.Display
	interfaceObj = structs.NewAnimal("Dog", 2)
	var newPerson structs.Person
	newPerson = structs.NewPerson("Truman", 14, true)
	fmt.Println(interfaceObj.Print())
	fmt.Println(newPerson.Print())
	*/
	fmt.Println(mock.IsEven(0))

}
