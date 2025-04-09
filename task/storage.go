package task

import (
	"encoding/json"
	"fmt"
	"os"
)

const filename = "tasks.json"

// function to save task to file
func SaveTasks(tasks []Task) error {

	file, err := json.MarshalIndent(tasks, "", "  ") // configure json , 2nd params -> prefix, 3rd params -> indentation

	if err != nil {
		return fmt.Errorf("Error saving tasks: %w", err)
	}

	err = os.WriteFile(filename, file, 0644) // add task from slice to json

	if err != nil {
		return fmt.Errorf("Error writing to file: %w", err)
	}

	return nil

}

// function to load all tasks
func LoadTasks()([]Task, error) {

	file, err := os.ReadFile(filename)

	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		
		return nil, fmt.Errorf("Error loading tasks: %w", err)
	}

	var tasks []Task

	err = json.Unmarshal(file, &tasks) // loads all tasks from json to slice

	if err != nil {
		return nil, fmt.Errorf("Error parsing tasks: %w", err)
	}

	return tasks, nil
}