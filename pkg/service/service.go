package service

import (
	"fmt"

	"github.com/jorgefg4/todolist/pkg/database"
	"github.com/jorgefg4/todolist/pkg/server"
)

// type Service struct {
// 	ID   string `boil:"ID" json:"ID"`
// 	Name string `boil:"name" json:"name,omitempty"`
// 	//Description string `json:"description,omitempty"`
// 	Check string `boil:"check" json:"check,omitempty"`
// }

// type Service interface {
// 	// NewService
// 	GetConnection() *sql.DB
// 	GetAllTasks() (map[int]task.Task, error)
// }

func NewService() server.Server {
	//var tasks map[int]*task.Task
	//tasks = data.Tasks

	tasks, _ := database.GetAllTasks()

	for _, task := range tasks {
		fmt.Println(task.Name)
	}

	//Llamo al package "server" para crear un nuevo router
	repo := database.NewGopherRepository(tasks)
	s := server.New(repo)

	return s
}
