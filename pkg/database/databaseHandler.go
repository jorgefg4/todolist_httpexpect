package database

import (
	"database/sql"

	"github.com/jorgefg4/todolist/pkg/task"
)

// Interface for accesses to a database
type DatabaseHandler interface {
	GetConnection() (*sql.DB, error)
	GetAllTasks() (map[int]*task.Task, error)
	CreateNewTask(name string) error
}
