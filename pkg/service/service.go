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
// 	Check_valid string `boil:"check_valid" json:"check_valid,omitempty"`
// }

// type Service interface {
// 	// NewService
// 	GetConnection() *sql.DB
// 	GetAllTasks() (map[int]task.Task, error)
// }

type Service struct {
	DB database.DatabaseHandler
}

func NewService(DB database.DatabaseHandler) *Service {
	return &Service{
		DB: DB,
	}
}

// TODO Gestionar errores
func (svc *Service) NewServer() (server.Server, error) {
	//var tasks map[int]*task.Task
	//tasks = data.Tasks

	err := svc.DB.GetConnection()
	if err != nil {
		return nil, err
	}

	/*
		// Pruebas de conexion
		svc.DB.CreateNewTask("Regar mis cactuses")
		if err != nil {
			fmt.Println(err)
		}
		svc.DB.CreateNewTask("Regar mis cactuses de nuevo")
		if err != nil {
			fmt.Println(err)
		}
	*/

	tasks, err := svc.DB.GetAllTasks()
	if err != nil {
		return nil, err
	}

	for _, task := range tasks {
		fmt.Println(task.Name)
	}

	//Llamo al package "server" para crear un nuevo router
	repo := database.NewTaskRepository(tasks, svc.DB)
	s := server.New(repo)

	return s, err
}
