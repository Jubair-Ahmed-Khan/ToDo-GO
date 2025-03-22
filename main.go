package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// task struct
type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

var tasks []Task
var nextID = 1

const filename = "tasks.json"

// ANSI escape codes for coloring text
const (
	ColorRed     = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorYellow  = "\033[33m"
	ColorBlue    = "\033[34m"
	ColorMagenta = "\033[35m"
	ColorReset   = "\033[0m"
)

// function to colorize text
func colorize(text string, colorCode string) string {
	return colorCode + text + ColorReset
}

func main() {

	// loading tasks from json file
	loadTasks()

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Println("\n" + colorize("#### To-Do List Application ####", ColorYellow))

		fmt.Println("\n" + colorize("1. View Task", ColorBlue))
		fmt.Println(colorize("2. Add Task", ColorBlue))
		fmt.Println(colorize("3. Edit Task", ColorBlue))
		fmt.Println(colorize("4. Toggle Task Completion", ColorBlue))

		// user choice input
		fmt.Print("\nChoose an option: ")

		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			viewTasks()
		case "2":
			fmt.Print("Enter task name: ")
			taskName, _ := reader.ReadString('\n')
			taskName = strings.TrimSpace(taskName)
			addTask(taskName)
		case "3":
			fmt.Print("Enter task ID to edit: ")
			taskIDStr, _ := reader.ReadString('\n')
			taskIDStr = strings.TrimSpace(taskIDStr)
			taskID, err := strconv.Atoi(taskIDStr)

			if err != nil {
				fmt.Println(colorize("Invalid task ID.", ColorRed))
				break
			}

			fmt.Print("Enter new task name: ")
			newName, _ := reader.ReadString('\n')
			newName = strings.TrimSpace(newName)
			editTask(taskID, newName)

		case "4":
			fmt.Print("Enter task ID to toggle completion: ")

			taskIDStr, _ := reader.ReadString('\n')
			taskIDStr = strings.TrimSpace(taskIDStr)
			taskID, err := strconv.Atoi(taskIDStr)

			if err != nil {
				fmt.Println(colorize("Invalid task ID.", ColorRed))
				break
			}

			toggleTaskCompletion(taskID)
			
		default:
			fmt.Println(colorize("\nInvalid option. Please try again.", ColorRed))
		}
	}
}

// function that loads all tasks
func loadTasks() {

	file, err := os.ReadFile(filename)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("\nNo existing tasks found. Starting fresh.")
			return
		}
		fmt.Println("Error loading tasks:", err)
		return
	}

	err = json.Unmarshal(file, &tasks) // loads all tasks from json to slice

	if err != nil {
		fmt.Println("Error parsing tasks:", err)
		return
	}

	// sets new task ID
	if len(tasks) > 0 {
		nextID = tasks[len(tasks)-1].ID + 1
	}
}

// function to view all tasks
func viewTasks() {

	// no task available
	if len(tasks) == 0 {
		fmt.Println("\n" + colorize("No tasks available.", ColorRed))
		return
	}

	// display all tasks
	fmt.Println("\n" + colorize("Tasks:", ColorMagenta))

	for _, task := range tasks {

		status := "Incomplete"

		if task.Done {
			status = "Complete"
		}

		fmt.Printf("%d: %s ", task.ID, task.Name)

		if status == "Complete" {
			fmt.Println(colorize("[Complete]", ColorGreen))
		} else {
			fmt.Println(colorize("[Incomplete]", ColorRed))
		}
	}
}

// function to add a task
func addTask(name string) {
	task := Task{
		ID:   nextID,
		Name: name,
		Done: false,
	}
	tasks = append(tasks, task)
	nextID++
	saveTasks()
	fmt.Println(colorize("\nTask added successfully!", ColorGreen))
}

// function to edit a task
func editTask(id int, newName string) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Name = newName
			saveTasks()
			fmt.Println(colorize("\nTask updated successfully!", ColorGreen))
			return
		}
	}
	fmt.Println(colorize("\nTask not found", ColorRed))
}

// function to toggle task status
func toggleTaskCompletion(id int) {

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = !tasks[i].Done
			saveTasks()
			fmt.Println(colorize("\nTask completion toggled successfully!", ColorGreen))
			return
		}
	}

	fmt.Println(colorize("\nTask not found", ColorRed))

}

// function to save task to file
func saveTasks() {

	file, err := json.MarshalIndent(tasks, "", "  ") // configure json , 2nd params -> prefix, 3rd params -> indentation

	if err != nil {
		fmt.Println(colorize("Error saving tasks:", ColorRed), err)
		return
	}

	err = os.WriteFile(filename, file, 0644) // add task from slice to json

	if err != nil {
		fmt.Println(colorize("Error writing to file:", ColorRed), err)
	}
}
