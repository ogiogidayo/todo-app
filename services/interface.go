package services

import (
	"context"

	"github.com/ogiogidayo/todo-app/database"
	"github.com/ogiogidayo/todo-app/domain"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . TaskAdder TaskLister
type TaskAdder interface {
	AddTask(ctx context.Context, db database.Execer, t *domain.Task) error
}

type TaskLister interface {
	ListTasks(ctx context.Context, db database.Queryer) (domain.Tasks, error)
}

type UserRegister interface {
	RegisterUser(ctx context.Context, db database.Execer, u *domain.User) error
}
