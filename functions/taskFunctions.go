package functions

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/DURAIRAJ17996/tutorial-go-basics/structs"
)

func AddTask(newTask structs.Task) bool {
	file, err := os.OpenFile("tasks.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return errorOccured()
	}
	newTask.Name = newTask.Name + ","
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("%+v\n", newTask))
	if err != nil {
		return errorOccured()
	}
	return true
}

func ListTasks() []structs.Task {
	tasks := []structs.Task{}
	data, err := ioutil.ReadFile("tasks.txt")

	if err != nil {
		fmt.Println("Error Occured")
		return nil
	}

	for _, line := range strings.Split(string(data), "\n") {
		if line == "" {
			return tasks
		}
		fields := strings.Split(line, ",")
		task := structs.Task{Name: fields[0], Done: fields[1] == "true"}

		tasks = append(tasks, task)
	}
	return tasks
}

func errorOccured() bool {
	fmt.Println("Error Occured")
	return false
}
