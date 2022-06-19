package database

import (
	"context"
	"database/sql"

	"github.com/jorgefg4/todolist/pkg/task"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/jorgefg4/todolist/models"
	//. "github.com/volatiletech/sqlboiler/v4/queries/qm"

	_ "github.com/lib/pq"
)

const conString string = "postgresql://postgres:gatomagico4444@localhost/postgres?sslmode=disable"

var db *sql.DB
var ctx context.Context

// Stablish a connection with the database
func GetConnection() (*sql.DB, error) {
	var err error
	db, err = sql.Open("postgres", conString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	ctx = context.Background()

	return db, err
}

// Retrieves all tasks from the database and returns a map with the tasks formatted
// for the application to use (type of task,Task)
func GetAllTasks() (map[string]*task.Task, error) {
	tasks, err := models.Tasks().All(ctx, db)
	if err != nil {
		return nil, err
	}

	tasksMap := make(map[string]*task.Task)

	for _, databaseTask := range tasks {
		id := databaseTask.ID
		name := databaseTask.Name
		check_valid := databaseTask.CheckValid

		newTask := task.Task{
			ID:    id,
			Name:  name,
			Check: check_valid,
		}

		tasksMap[id] = &newTask
	}

	return tasksMap, err
}

// Create a new task in the database
// Should use upsert instead of insert in the future
func CreateNewTask(id string, name string, check bool) error {
	var newTask models.Task

	newTask.ID = id
	newTask.Name = name
	newTask.CheckValid = check

	err := newTask.Insert(ctx, db, boil.Infer())

	return err
}

/*
func DeleteTask(id string) error {
	return error
}

func ModifyTask(id string, name string, check bool) error {
	return error
}
*/
