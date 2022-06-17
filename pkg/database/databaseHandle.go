package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jorgefg4/todolist/models"
	//"github.com/volatiletech/sqlboiler/boil"
	//. "github.com/volatiletech/sqlboiler/v4/queries/qm"

	_ "github.com/lib/pq"
)

const conString string = "postgresql://postgres:gatomagico4444@localhost/postgres?sslmode=disable"

var db *sql.DB
var ctx context.Context

func GetConnection() *sql.DB {
	var err error
	db, err = sql.Open("postgres", conString)
	if err != nil {
		fmt.Println(err)
	}

	//boil.SetDB(db)

	ctx = context.Background()

	return db
}

func GetTasks() {
	tasks, err := models.Tasks().All(ctx, db)
	if err != nil {
		fmt.Println(err)
	}

	for _, task := range tasks {
		fmt.Println(task)
	}
}
