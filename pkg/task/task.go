package task

//Se incluyen tags para tranformar el struct desde json o hacia json
type Task struct {
	ID    string `json:"ID"`
	Name  string `json:"name,omitempty"`
	Check bool   `json:"check,omitempty"`
}

// Repository provides access to the task storage
type TaskRepository interface {
	// CreateTask saves a given task
	CreateTask(g *Task) error
	// FetchTasks return all tasks saved in storage
	FetchTasks() ([]*Task, error)
	// DeleteTask remove task with given ID
	DeleteTask(ID string) error
	// UpdateTask modify task with given ID and given new data
	// UpdateTask(ID string, g *Task) error
	UpdateTask(ID string) (int, error)
	// FetchTaskByID returns the task with given ID
	FetchTaskByID(ID string) (*Task, error)
}
