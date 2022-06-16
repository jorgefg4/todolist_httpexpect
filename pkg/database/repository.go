package database

import (
	"fmt"
	"sync"

	task "github.com/jorgefg4/todolist/pkg/task"
)

type taskRepository struct {
	mtx   sync.RWMutex
	tasks map[string]*task.Task
}

func NewTaskRepository(tasks map[string]*task.Task) task.TaskRepository {
	if tasks == nil {
		tasks = make(map[string]*task.Task)
	}

	return &taskRepository{
		tasks: tasks,
	}
}

func (r *taskRepository) CreateTask(g *task.Task) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if err := r.checkIfExists(g.ID); err != nil {
		return err
	}
	g.Check = false
	r.tasks[g.ID] = g

	return nil
}

func (r *taskRepository) FetchTasks() ([]*task.Task, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	values := make([]*task.Task, 0, len(r.tasks))
	for _, value := range r.tasks {
		values = append(values, value)
	}
	return values, nil
}

func (r *taskRepository) DeleteTask(ID string) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	delete(r.tasks, ID)

	return nil
}

func (r *taskRepository) UpdateTask(ID string) (int, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	for _, v := range r.tasks {
		if v.ID == ID {
			r.tasks[ID].Check = true
			return 0, nil
		}
	}
	return 1, nil

}

func (r *taskRepository) FetchTaskByID(ID string) (*task.Task, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	for _, v := range r.tasks {
		if v.ID == ID {
			return v, nil
		}
	}

	return nil, fmt.Errorf("The ID %s doesn't exist", ID)
}

func (r *taskRepository) checkIfExists(ID string) error {
	for _, v := range r.tasks {
		if v.ID == ID {
			return fmt.Errorf("The task %s is already exist", ID)
		}
	}

	return nil
}
