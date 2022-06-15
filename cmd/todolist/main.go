package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jorgefg4/todolist/pkg/data"
	"github.com/jorgefg4/todolist/pkg/database"
	"github.com/jorgefg4/todolist/pkg/server"
	"github.com/jorgefg4/todolist/pkg/task"

	//_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
)

func main() {
	fmt.Printf("Hello world!\n")

	// // 	const create string = `
	// //   CREATE TABLE IF NOT EXISTS activities (
	// //   id INTEGER NOT NULL PRIMARY KEY,
	// //   time DATETIME NOT NULL,
	// //   description TEXT
	// //   );`

	// // Get a handle to the SQLite database, using mattn/go-sqlite3
	// db, err := sql.Open("sqlite3", "./test.sqlite3")
	// if err != nil {
	// 	panic(err)
	// }
	// // Configure SQLBoiler to use the sqlite database
	// boil.SetDB(db)
	// ctx := context.Background()
	// books, _ := models.Books().All(ctx, db)

	// // t := &models.Task{Name: "prueba"}
	// // err = t.Insert(context.Background(), db, boil.Infer())
	// // fmt.Println("task name:", t.Name)

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
