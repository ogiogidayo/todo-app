package handler

import (
	"context"

	"github.com/ogiogidayo/todo-app/domain"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . ListTaskService AddTaskService RegisterUserService
type ListTaskService interface {
	ListTasks(ctx context.Context) (domain.Tasks, error)
}

type AddTaskService interface {
	AddTask(ctx context.Context, title string) (*domain.Task, error)
}

type RegisterUserService interface {
	RegisterUser(ctx context.Context, name, password, role string) (*domain.User, error)
}
