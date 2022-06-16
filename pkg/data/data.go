package data

import (
	task "github.com/jorgefg4/todolist/pkg/task"
)

var Tasks = map[string]*task.Task{
	"01D3XZ3ZHCP3KG9VT4FGAD8KDR": &task.Task{
		ID:    "01D3XZ3ZHCP3KG9VT4FGAD8KDR",
		Name:  "Tarea1",
		Check: "no",
	},
	"01D3XZ7CN92AKS9HAPSZ4D5DP9": &task.Task{
		ID:    "01D3XZ7CN92AKS9HAPSZ4D5DP9",
		Name:  "Tarea2",
		Check: "si",
	},
	"01D3XZ89NFJZ9QT2DHVD462AC2": &task.Task{
		ID:    "01D3XZ89NFJZ9QT2DHVD462AC2",
		Name:  "Tarea3",
		Check: "si",
	},
	"01D3XZ8JXHTDA6XY05EVJVE9Z2": &task.Task{
		ID:    "01D3XZ8JXHTDA6XY05EVJVE9Z2",
		Name:  "Tarea4",
		Check: "no",
	},
}
