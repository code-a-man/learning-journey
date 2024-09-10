package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Progress uint8  `json:"progress"`
}

var progressStatus = []string{"todo", "in-progress", "done"}

var filename = "tasks.json"

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Usage: task-tracker <command> <args>")
		return
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		createFile()
	}

	tasks, err := readTasks(filename)

	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	command := args[0]
	arg1 := ""
	if len(args) > 1 {
		arg1 = args[1]
	}
	arg2 := ""
	if len(args) > 2 {
		arg2 = args[2]
	}
	switch command {
	case "add":
		tasks = append(tasks, Task{ID: len(tasks) + 1, Name: arg1, Progress: 0})
		err := updateFile(tasks)
		if err != nil {
			fmt.Println("Error updating file: ", err)
			return
		}
		fmt.Println("Task added successfully (ID: ", len(tasks), ")")
	case "list":
		listTasks(tasks, arg1)
	case "update":
		id := 0
		fmt.Sscanf(arg1, "%d", &id)
		if id == 0 || id > len(tasks) {
			fmt.Println("Invalid ID")
			return
		}
		tasks[id-1].Name = arg2
		err := updateFile(tasks)
		if err != nil {
			fmt.Println("Error updating file: ", err)
			return
		}
		fmt.Println("Task updated successfully")
	case "mark":
		id := 0
		fmt.Sscanf(arg1, "%d", &id)
		if id == 0 || id > len(tasks) {
			fmt.Println("Invalid ID")
			return
		}
		status := -1
		for i, s := range progressStatus {
			if s == arg2 {
				status = i
				break
			}
		}
		if status == -1 {
			fmt.Println("Invalid status")
			fmt.Println("Available statuses: ", progressStatus)
			return
		}
		tasks[id-1].Progress = uint8(status)
		err := updateFile(tasks)
		if err != nil {
			fmt.Println("Error updating file: ", err)
			return
		}
		fmt.Println("Task marked successfully")
	case "delete":
		id := 0
		fmt.Sscanf(arg1, "%d", &id)
		if id == 0 || id > len(tasks) {
			fmt.Println("Invalid ID")
			return
		}
		tasks = append(tasks[:id-1], tasks[id:]...)

		updateFile(tasks)

	default:
		fmt.Println("Invalid command")
		fmt.Println("Usage: task-tracker add <task>")
	}

}

func readTasks(filename string) ([]Task, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	var tasks []Task
	json.Unmarshal(byteValue, &tasks)

	return tasks, nil

}

func createFile() error {
	file, err := os.Create("tasks.json")
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return err
	}
	defer file.Close()

	_, err = file.WriteString("[]")
	if err != nil {
		fmt.Println("Error writing to file: ", err)
		return err
	}

	fmt.Println("File created successfully")
	return nil
}

func updateFile(tasks []Task) error {
	file, err := os.OpenFile(filename, os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return err
	}
	defer file.Close()

	jsonTasks, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println("Error marshalling tasks: ", err)
		return err
	}

	_, err = file.Write(jsonTasks)
	if err != nil {
		fmt.Println("Error writing to file")
		return err
	}

	return nil
}

func getStatusString(done uint8) string {
	if int(done) < len(progressStatus) {
		return progressStatus[done]
	} else {
		return "invalid"
	}
}

func listTasks(tasks []Task, filter string) {
	fmt.Println("ID\tTask\tStatus")
	filteredTasks := []Task{}
	if filter != "" {
		for _, task := range tasks {
			if getStatusString(task.Progress) == filter {
				filteredTasks = append(filteredTasks, task)
			}
		}
	} else {
		filteredTasks = tasks
	}

	for _, task := range filteredTasks {
		fmt.Println(task.ID, "\t", task.Name, "\t", getStatusString(task.Progress))
	}
}
