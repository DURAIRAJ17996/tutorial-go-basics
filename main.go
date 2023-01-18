package main

import (
	"fmt"
	"os"

	"github.com/DURAIRAJ17996/tutorial-go-basics/functions"
	"github.com/DURAIRAJ17996/tutorial-go-basics/structs"
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
	//fmt.Println(mock.IsEven(0))
	var taskList []structs.Task
	taskList = make([]structs.Task, 0)
	var taskName string
	cmd := os.Args[1]
	if len(os.Args) > 1 {
		taskName = os.Args[2]
	}

	switch cmd {
	case "add":
		newTask := structs.Task{taskName, false}
		var saveFlag = functions.AddTask(newTask)
		fmt.Println("Save task: ", saveFlag)
		os.Exit(1)
	case "list":
		fmt.Println("Fetching tasks")
		fmt.Println(functions.ListTasks())
		os.Exit(1)
	}

	fmt.Println(taskList)

}
