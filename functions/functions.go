package functions

import "fmt"

// Function with multiple returns
func FunctionWithTwoReturn(index int) (int, string) {
	fmt.Println("Inside functions.FunctionWithTwoReturn")
	fmt.Println("Index: ", index)
	varMap := make(map[int]string)
	varMap[1] = "Samson"
	varMap[2] = "Ruturaj"
	varMap[3] = "Gill"

	fmt.Println("Value at index: ", varMap[index])
	return index, varMap[index]
}

// Function with variable input arguments
func FunctionWithVariableArgs(stringList ...string) {
	for _, varCurrent := range stringList {
		fmt.Println("Current string value: ", varCurrent)
	}
}

// Anonymous function
func FunctionWithAnonymousFunc() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
