package services

import (
	"context"
	"fmt"

	"github.com/ogiogidayo/todo-app/database"
	"github.com/ogiogidayo/todo-app/domain"
)

type AddTask struct {
	DB   database.Execer
	Repo TaskAdder
}

func (a *AddTask) AddTask(ctx context.Context, title string) (*domain.Task, error) {
	t := &domain.Task{
		Title:  title,
		Status: domain.TaskStatusTodo,
	}
	err := a.Repo.AddTask(ctx, a.DB, t)
	if err != nil {
		return nil, fmt.Errorf("failed to register: %w", err)
	}
	return t, nil
}
