package main

import (
	// "ToDoList/data"
	// "ToDoList/database"
	// "ToDoList/server"
	// "ToDoList/task"
	"fmt"
	"log"
	"net/http"

	"github.com/jorgefg4/todolist/pkg/database"
	"github.com/jorgefg4/todolist/pkg/server"
	"github.com/rs/cors"
)

func main() {
	_, err := database.GetConnection()
	if err != nil {
		fmt.Println(err)
	}

	tasks, err := database.GetAllTasks()
	if err != nil {
		fmt.Println(err)
	}

	// Pruebas de conexion
	database.CreateNewTask("Regar mis cactuses")
	if err != nil {
		fmt.Println(err)
	}
	database.CreateNewTask("Regar mis cactuses de nuevo")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Server running\n")

	//if *withData {
	//tasks = data.Tasks

	//Llamo al package "server" para crear un nuevo router
	repo := database.NewGopherRepository(tasks)
	s := server.New(repo)

	//Cabeceras CORS:
	handler := cors.New(cors.Options{AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"}}).Handler(s.Router())
	log.Fatal(http.ListenAndServe(":8000", handler))

	//http.ListenAndServe(":8000", s.Router()) //Se pone a escuchar en el puerto TCP 8000 de localhost y llama al handler

}
