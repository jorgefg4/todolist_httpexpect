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
	// CreateGopher saves a given gopher
	CreateGopher(g *Task) error
	// FetchGophers return all gophers saved in storage
	FetchGophers() ([]*Task, error)
	// DeleteGopher remove gopher with given ID
	DeleteGopher(ID string) error
	// UpdateGopher modify gopher with given ID and given new data
	// UpdateGopher(ID string, g *Task) error
	UpdateGopher(ID string) (int, error)
	// FetchGopherByID returns the gopher with given ID
	FetchGopherByID(ID string) (*Task, error)
}
