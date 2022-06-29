package task

//Se incluyen tags para tranformar el struct desde json o hacia json
type Task struct {
	ID   int    `json:"ID"`
	Name string `json:"name,omitempty"`
	//Description string `json:"description,omitempty"`
	Check_valid bool `json:"check_valid,omitempty"`
}

// Repository provides access to the task storage
type TaskRepository interface {
	// CreateGopher saves a given gopher
	CreateTask(g *Task) error
	// FetchGophers return all gophers saved in storage
	FetchTasks() ([]*Task, error)
	// DeleteGopher remove gopher with given ID
	DeleteTask(ID int) error
	// UpdateGopher modify gopher with given ID and given new data
	// UpdateGopher(ID int, g *Task) error
	UpdateTask(ID int) (int, error)
	// FetchGopherByID returns the gopher with given ID
	FetchTaskByID(ID int) (*Task, error)
}
