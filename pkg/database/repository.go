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

func NewGopherRepository(tasks map[string]*task.Task) task.TaskRepository {
	if tasks == nil {
		tasks = make(map[string]*task.Task)
	}

	return &taskRepository{
		tasks: tasks,
	}
}

func (r *taskRepository) CreateGopher(g *task.Task) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if err := r.checkIfExists(g.ID); err != nil {
		return err
	}
	r.tasks[g.ID] = g
	return nil
}

func (r *taskRepository) FetchGophers() ([]*task.Task, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	values := make([]*task.Task, 0, len(r.tasks))
	for _, value := range r.tasks {
		values = append(values, value)
	}
	return values, nil
}

func (r *taskRepository) DeleteGopher(ID string) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	delete(r.tasks, ID)

	return nil
}

func (r *taskRepository) UpdateGopher(ID string, g *task.Task) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	g.Check = "si"
	r.tasks[ID] = g
	return nil
}

func (r *taskRepository) FetchGopherByID(ID string) (*task.Task, error) {
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
