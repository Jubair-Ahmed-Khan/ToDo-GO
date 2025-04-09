package userView

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Jubair-Ahmed-Khan/ToDoList/task"
	"github.com/Jubair-Ahmed-Khan/ToDoList/utils"
)

type CLI struct {
	tasks  []task.Task
	nextID int
	reader *bufio.Reader
}

// function to create instance of the program
func NewCLI() *CLI {
	tasks, err := task.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		os.Exit(1)
	}

	nextID := 1
	if len(tasks) > 0 {
		nextID = tasks[len(tasks)-1].ID + 1
	}

	return &CLI{
		tasks:  tasks,
		nextID: nextID,
		reader: bufio.NewReader(os.Stdin),
	}
}

// function to run the program
func (c *CLI) Run() {
	for {

		c.displayMenu()
		option := c.getInput("\nChoose an option: ")

		switch option {
		case "1":
			c.viewTasks()
		case "2":
			c.addTask()
		default:
			fmt.Println(utils.Colorize("\nInvalid option. Please try again.", utils.ColorRed))
		}
	}
}

// function to display user menu
func (c *CLI) displayMenu() {
	fmt.Println("\n" + utils.Colorize("#### To-Do List Application ####", utils.ColorYellow))

	fmt.Println("\n" + utils.Colorize("1. View Task", utils.ColorBlue))
	fmt.Println(utils.Colorize("2. Add Task", utils.ColorBlue))
}

// function to get user input
func (c *CLI) getInput(prompt string) string {
	fmt.Print(prompt)
	input, _ := c.reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// function to add a task
func (c *CLI) addTask() {
	name := c.getInput("Enter task name: ")
	task := task.NewTask(c.nextID, name)
	c.tasks = append(c.tasks, task)
	c.nextID++
	c.saveTasks()
	fmt.Println(utils.Colorize("\nTask added successfully!", utils.ColorGreen))
}

// function to view tasks
func (c *CLI) viewTasks() {

	// no task available
	if len(c.tasks) == 0 {
		fmt.Println("\n" + utils.Colorize("No tasks available.", utils.ColorRed))
		return
	}

	// display all tasks
	fmt.Println("\n" + utils.Colorize("Tasks:", utils.ColorMagenta))

	for _, task := range c.tasks {

		status := "Incomplete"

		if task.Done {
			status = "Complete"
		}

		fmt.Printf("%d: %s ", task.ID, task.Name)

		if status == "Complete" {
			fmt.Println(utils.Colorize("[Complete]", utils.ColorGreen))
		} else {
			fmt.Println(utils.Colorize("[Incomplete]", utils.ColorRed))
		}
	}
}

func (c *CLI) saveTasks() {
	err := task.SaveTasks(c.tasks)
	if err != nil {
		fmt.Println(utils.Colorize("Error saving tasks:", utils.ColorRed), err)
	}
}
