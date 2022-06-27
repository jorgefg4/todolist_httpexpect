package database

import (
	"fmt"
	"testing"

	task "github.com/jorgefg4/todolist/pkg/task"
	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {

	_, err := GetConnection()
	if err != nil {
		fmt.Println(err)
	}

	tasks, err := GetAllTasks()
	if err != nil {
		fmt.Println(err)
	}

	repo := NewTaskRepository(tasks)

	t1 := task.Task{ID: 1, Name: "tarea de prueba"}
	repo.CreateTask(&t1)

	repo1, _ := repo.FetchTasks()
	for _, value := range repo1 {
		if value.Name == "tarea de prueba" {
			assert.Equal(t, value.Name, t1.Name, "The two tasks should be the same.")
		}
	}
}
