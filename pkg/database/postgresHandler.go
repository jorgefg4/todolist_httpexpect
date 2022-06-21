package database

import (
	"context"
	"database/sql"

	task "github.com/jorgefg4/todolist/pkg/task"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/jorgefg4/todolist/models"
	//. "github.com/volatiletech/sqlboiler/v4/queries/qm"

	_ "github.com/lib/pq"
)

// Type to access a Postgresql database
// Implements DatabaseHandler
type PostgresHandler struct {
}

const conString string = "postgresql://postgres:gatomagico4444@localhost/postgres?sslmode=disable"

var db *sql.DB
var ctx context.Context

// Stablish a connection with the database
func (handler *PostgresHandler) GetConnection() (*sql.DB, error) {
	var err error
	db, err = sql.Open("postgres", conString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// TODO definir el contexto
	ctx = context.Background()

	return db, err
}

// Retrieves all tasks from the database and returns a map with the tasks formatted
// for the application to use (type of task,Task)
func (handler *PostgresHandler) GetAllTasks() (map[int]*task.Task, error) {
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

	return tasksMap, err
}

// Create a new task in the database
// Should use upsert instead of insert in the future
func (handler *PostgresHandler) CreateNewTask(name string) error {
	var newTask models.Task

	newTask.Name = name

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
