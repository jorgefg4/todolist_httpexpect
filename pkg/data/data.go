package data

import (
	task "github.com/jorgefg4/todolist/pkg/task"
)

var Tasks = map[int]*task.Task{
	1: &task.Task{
		ID:    1,
		Name:  "Tarea1",
		Check: false,
	},
	2: &task.Task{
		ID:    2,
		Name:  "Tarea2",
		Check: true,
	},
	3: &task.Task{
		ID:    3,
		Name:  "Tarea3",
		Check: true,
	},
	4: &task.Task{
		ID:    4,
		Name:  "Tarea4",
		Check: false,
	},
}
