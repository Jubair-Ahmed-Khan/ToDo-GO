package task

// task struct
type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

// return a new task
func NewTask(id int, name string) Task {
	return Task{
		ID:   id,
		Name: name,
		Done: false,
	}
}

// toggle task completion status
func (t *Task) ToggleCompletion() {
	t.Done = !t.Done
}
