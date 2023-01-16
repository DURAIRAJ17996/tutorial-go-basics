package datastructure

import (
	"fmt"
)

/*
 1. Syntax to declare and initialize slice
    SYNTAX: varaiableName := make([]data-type, lengthOfArray)
*/
func TestSlice() []string {
	fmt.Println("Inside datastructure.TestSlice")
	varSlice := make([]string, 3)
	fmt.Println("Length of  varSlice: ", len(varSlice))
	// append to add something at the end
	varSlice = append(varSlice, "Dhoni")
	varSlice = append(varSlice, "Sachin")
	// setting a value at an index
	varSlice[3] = "Virat"
	return varSlice
}
