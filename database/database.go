package database

import (
	"errors"

	"github.com/ogiogidayo/todo-app/domain"
)

var (
	Tasks = &TaskStore{Tasks: map[domain.TaskID]*domain.Task{}}

	ErrNotFound = errors.New("not found")
)

type TaskStore struct {
	LastID domain.TaskID
	Tasks  map[domain.TaskID]*domain.Task
}

func (ts *TaskStore) Add(t *domain.Task) (domain.TaskID, error) {
	ts.LastID++
	t.ID = ts.LastID
	ts.Tasks[t.ID] = t
	return t.ID, nil
}

func (ts *TaskStore) Get(id domain.TaskID) (*domain.Task, error) {
	if ts, ok := ts.Tasks[id]; ok {
		return ts, nil
	}
	return nil, ErrNotFound
}

func (ts *TaskStore) All() domain.Tasks {
	tasks := make([]*domain.Task, len(ts.Tasks))
	for i, t := range ts.Tasks {
		tasks[i-1] = t
	}
	return tasks
}
