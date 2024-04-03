package services

import (
	"context"
	"fmt"
	"github.com/ogiogidayo/todo-app/database"
	"github.com/ogiogidayo/todo-app/entity"
)

type AddTask struct {
	DB   database.Execer
	Repo TaskAdder
}

func (a *AddTask) AddTask(ctx context.Context, title string) (*entity.Task, error) {
	t := &entity.Task{
		Title:  title,
		Status: entity.TaskStatusTodo,
	}
	err := a.Repo.AddTask(ctx, a.DB, t)
	if err != nil {
		return nil, fmt.Errorf("failed to register: %w", err)
	}
	return t, nil
}

type ListTask struct {
	DB   database.Queryer
	Repo TaskLister
}

func (l *ListTask) ListTask(ctx context.Context) (entity.Tasks, error) {
	ts, err := l.Repo.ListTask(ctx, l.DB)
	if err != nil {
		return nil, fmt.Errorf("faild to list: %w", err)
	}
	return ts, nil
}
