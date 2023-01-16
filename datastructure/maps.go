package datastructure

import "fmt"

/*
Syntax to create map:
1. Using make and map
 SYNTAX: variableName := make(map[key-type]value-type)
2. Using just map
 SYNTAX: map[key-type]value-type{k1:v1, k2:v2}

*/
func TestMap() map[int]string {
	fmt.Println("Inside datastructure.TestMap")
	varMap := make(map[int]string)
	varMap[1] = "Sachin"
	varMap[2] = "Dhoni"
	varMap[3] = "Virat"
	return varMap
}

func TestMapWithRange() {
	fmt.Println("Inside datastructure.TestMapWithRange")
	varMap := map[int]string{1: "Carl", 2: "Sagan", 3: "Cosmos"}

	// Different ways to loop on Map using for and range
	// _ (blank identifier) can be used when we don't need the index
	for _, varString := range varMap {
		fmt.Println("Key, value ", varString)
	}
	// k, v - Key and value of map
	for k, v := range varMap {
		fmt.Println("Key, value ", k, v)
	}
	// k - Key of the map, (we cannot use value like this)
	for k := range varMap {
		fmt.Println("Key ", k)
	}
}
