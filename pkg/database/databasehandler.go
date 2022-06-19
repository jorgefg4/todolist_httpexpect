package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jorgefg4/todolist/models"
	task "github.com/jorgefg4/todolist/pkg/task"
)

const conString string = "postgresql://postgres:gatomagico4444@localhost/postgres?sslmode=disable"

var db *sql.DB
var ctx context.Context

// Stablish a connection with the database
func GetConnection() *sql.DB {
	var err error
	db, err = sql.Open("postgres", conString)
	if err != nil {
		fmt.Println(err)
	}
	// Configure SQLBoiler to use the sqlite database
	//boil.SetDB(db)
	ctx = context.Background()

	return db
}

// Retrieves all tasks from the database and returns a map with the tasks formatted
// for the application to use (type of task,Task)
func GetAllTasks() (map[int]*task.Task, error) {
	tasks, err := models.Tasks().All(ctx, db)
	if err != nil {
		return nil, err
	}

	tasksMap := make(map[int]*task.Task)

	for _, databaseTask := range tasks {
		id := databaseTask.ID
		name := databaseTask.Name
		check := databaseTask.Check

		newTask := task.Task{
			ID:    id,
			Name:  name,
			Check: check,
		}

		tasksMap[id] = &newTask
	}

	return tasksMap, nil
}
