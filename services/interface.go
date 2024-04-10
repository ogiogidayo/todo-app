package services

import (
	"context"

	"github.com/ogiogidayo/todo-app/database"
	"github.com/ogiogidayo/todo-app/domain"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . TaskAdder TaskLister UserRegister UserGetter TokenGenerator
type TaskAdder interface {
	AddTask(ctx context.Context, db database.Execer, t *domain.Task) error
}
type TaskLister interface {
	ListTasks(ctx context.Context, db database.Queryer, id domain.UserID) (domain.Tasks, error)
}
type UserRegister interface {
	RegisterUser(ctx context.Context, db database.Execer, u *domain.User) error
}

type UserGetter interface {
	GetUser(ctx context.Context, db database.Queryer, name string) (*domain.User, error)
}

type TokenGenerator interface {
	GenerateToken(ctx context.Context, u domain.User) ([]byte, error)
}
