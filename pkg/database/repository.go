package database

import (
	"fmt"
	"sync"

	task "github.com/jorgefg4/todolist/pkg/task"
)

type taskRepository struct {
	mtx   sync.RWMutex
	tasks map[int]*task.Task
	db    DatabaseHandler
}

func NewTaskRepository(tasks map[int]*task.Task) task.TaskRepository {
	if tasks == nil {
		tasks = make(map[int]*task.Task)
	}

	return &taskRepository{
		tasks: tasks,
	}
}

// TODO gestion de errores
func (r *taskRepository) CreateTask(g *task.Task) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.db.CreateNewTask(g.Name) //llama a la función de añadir task de postgresHandler
	// if err := r.checkIfExists(g.ID); err != nil {
	// 	return err
	// }
	// g.Check = false
	// r.tasks[g.ID] = g

	return nil
}

func (r *taskRepository) FetchTasks() ([]*task.Task, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	//Obtengo tasks de la BD y actualizo el map de tasks del repository
	tasks, err := r.db.GetAllTasks()
	r.tasks = tasks
	if err != nil {
		fmt.Println(err)
	}

	values := make([]*task.Task, 0, len(r.tasks))
	for _, value := range r.tasks {
		values = append(values, value)
	}
	return values, nil
}

func (r *taskRepository) DeleteTask(ID int) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	r.db.DeleteTask(ID) //llamada a la funcion de borrar de postgresHandler
	//delete(r.tasks, ID)

	return nil
}

func (r *taskRepository) UpdateTask(ID int) (int, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	err := r.db.ModifyTask(ID)
	// for _, v := range r.tasks {
	// 	if v.ID == ID {
	// 		r.tasks[ID].Check = true
	// 		return 0, nil
	// 	}
	// }
	if err != nil {
		return 1, err
	} else {
		return 0, nil
	}

}

func (r *taskRepository) FetchTaskByID(ID int) (*task.Task, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	for _, v := range r.tasks {
		if v.ID == ID {
			return v, nil
		}
	}

	return nil, fmt.Errorf("The ID %d doesn't exist", ID)
}

func (r *taskRepository) checkIfExists(ID int) error {
	for _, v := range r.tasks {
		if v.ID == ID {
			return fmt.Errorf("The task %d is already exist", ID)
		}
	}

	return nil
}
