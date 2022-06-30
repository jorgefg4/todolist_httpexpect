package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	//_ "github.com/mattn/go-sqlite3"

	"github.com/jorgefg4/todolist/pkg/database"

	"github.com/jorgefg4/todolist/pkg/service"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func main() {
	fmt.Printf("Server running\n")

	// // Get a handle to the SQLite database
	// db, err := sql.Open("postgres", `dbname=postgres host=localhost user=postgres password=gatomagico4444`)
	// db, err := sql.Open("postgres", "postgresql://postgres:gatomagico4444@localhost/postgres?sslmode=disable")

	// if err != nil {
	// 	panic(err)
	// }

	// defer db.Close()
	// // Configure SQLBoiler to use the sqlite database
	// boil.SetDB(db)
	// ctx := context.Background()
	// tasks, _ := models.Tasks().All(ctx, db)

	// for _, task := range tasks {
	// 	fmt.Println(task.Name)
	// }

	// id := 3
	// task, err := models.FindTask(ctx, db, id)
	// if err != nil {
	// } // handle err
	// _, err = task.Delete(ctx, db)
	// if err != nil {
	// } // handle err

	// t := &models.Task{Name: "prueba29"}
	// err = t.Insert(context.Background(), db, boil.Infer())
	// fmt.Println("task name:", t.Name)
	// if err != nil {
	// 	panic(err)
	// }

	// t1 := &models.Task{Name: "prueba27"}
	// err = t1.Insert(context.Background(), db, boil.Infer())
	// fmt.Println("task name:", t1.Name)
	// if err != nil {
	// 	panic(err)
	// }

	// t2 := &models.Task{Name: "prueba28"}
	// err = t2.Insert(context.Background(), db, boil.Infer())
	// fmt.Println("task name:", t2.Name)
	// if err != nil {
	// 	panic(err)
	// }

	// var tasks map[string]*task.Task
	// //if *withData {
	// tasks = data.Tasks

	// //Llamo al package "server" para crear un nuevo router
	// repo := database.NewGopherRepository(tasks)
	// s := server.New(repo)

	// Declaramos aqui las variables
	var db *sql.DB
	var ctx context.Context = context.Background()

	//Llamo a la capa Service para crear el Server
	ph := database.NewPostgres(db, ctx)
	svc := service.NewService(ph)
	s, err := svc.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	//Cabeceras CORS:
	handler := cors.New(cors.Options{AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"}}).Handler(s.Router())

	log.Fatal(http.ListenAndServe(":8000", handler)) //Se pone a escuchar en el puerto TCP 8000 de localhost y llama al handler

	//http.ListenAndServe(":8000", s.Router()) //Se pone a escuchar en el puerto TCP 8000 de localhost y llama al handler

}
