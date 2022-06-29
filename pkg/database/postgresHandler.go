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
// TODO agregar como campos los datos de conexion sql (agregar aqui mas campos aparte de la db?)
// Tener en cuenta la posibilidad de agregar un logger
type PostgresHandler struct {
	DB  *sql.DB
	ctx context.Context // Deberia ir aqui...?
}

// TODO Plantear dejar en otro sitio mas adecuado (main?)
const conString string = "postgresql://postgres:gatomagico4444@db/postgres?sslmode=disable"

//var db *sql.DB
//var ctx context.Context

// TODO Como controlar errores aqui? Da problemas en el main donde se invoca
func NewPostgres(db *sql.DB, ctx context.Context) *PostgresHandler {
	return &PostgresHandler{
		DB:  db,
		ctx: ctx,
	}
}

// Stablish a connection with the database
func (handler *PostgresHandler) GetConnection() error {
	var err error
	handler.DB, err = sql.Open("postgres", conString)
	if err != nil {
		return err
	}

	err = handler.DB.Ping()
	if err != nil {
		return err
	}

	// TODO definir el contexto
	//ctx = context.Background()

	return err
}

// Retrieves all tasks from the database and returns a map with the tasks formatted
// for the application to use (type of task,Task)
func (handler *PostgresHandler) GetAllTasks() (map[int]*task.Task, error) {
	tasks, err := models.Tasks().All(handler.ctx, handler.DB)
	if err != nil {
		return nil, err
	}

	tasksMap := make(map[int]*task.Task)

	for _, databaseTask := range tasks {
		id := databaseTask.ID
		name := databaseTask.Name
		check_valid := databaseTask.CheckValid

		newTask := task.Task{
			ID:          id,
			Name:        name,
			Check_valid: check_valid,
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

	err := newTask.Insert(handler.ctx, handler.DB, boil.Infer())

	return err
}

// Delete a given task from the database
func (handler *PostgresHandler) DeleteTask(id int) error {
	task, err := models.FindTask(handler.ctx, handler.DB, id)

	_, err = task.Delete(handler.ctx, handler.DB)

	return err
}

// Modifies a given task from the database
func (handler *PostgresHandler) ModifyTask(id int) error {
	task, err := models.FindTask(handler.ctx, handler.DB, id)

	task.CheckValid = true

	_, err = task.Update(handler.ctx, handler.DB, boil.Infer())

	return err
}
