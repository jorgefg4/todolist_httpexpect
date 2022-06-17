package main

import (
	// "ToDoList/data"
	// "ToDoList/database"
	// "ToDoList/server"
	// "ToDoList/task"
	"fmt"
	"log"
	"net/http"

	"github.com/jorgefg4/todolist/pkg/data"
	"github.com/jorgefg4/todolist/pkg/database"
	"github.com/jorgefg4/todolist/pkg/server"
	"github.com/jorgefg4/todolist/pkg/task"
	"github.com/rs/cors"
)

func main() {
	database.GetConnection()
	database.GetTasks()

	fmt.Printf("Hello world!\n")

	var tasks map[string]*task.Task
	//if *withData {
	tasks = data.Tasks

	//Llamo al package "server" para crear un nuevo router
	repo := database.NewGopherRepository(tasks)
	s := server.New(repo)

	//Cabeceras CORS:
	handler := cors.New(cors.Options{AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"}}).Handler(s.Router())
	log.Fatal(http.ListenAndServe(":8000", handler))

	//http.ListenAndServe(":8000", s.Router()) //Se pone a escuchar en el puerto TCP 8000 de localhost y llama al handler

}
