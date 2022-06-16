package task

//Se incluyen tags para tranformar el struct desde json o hacia json
type Task struct {
	ID   string `json:"ID"`
	Name string `json:"name,omitempty"`
	//Description string `json:"description,omitempty"`
	Check string `json:"check,omitempty"`
}

// Repository provides access to the gopher storage
type TaskRepository interface {
	// CreateTask saves a given gopher
	CreateTask(g *Task) error
	// FetchTasks return all gophers saved in storage
	FetchTasks() ([]*Task, error)
	// DeleteTask remove gopher with given ID
	DeleteTask(ID string) error
	// UpdateTask modify gopher with given ID and given new data
	// UpdateTask(ID string, g *Task) error
	UpdateTask(ID string) (int, error)
	// FetchTaskByID returns the gopher with given ID
	FetchTaskByID(ID string) (*Task, error)
}
