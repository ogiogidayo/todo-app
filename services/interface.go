package services

import (
	"context"

	"github.com/ogiogidayo/todo-app/database"
	"github.com/ogiogidayo/todo-app/entity"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . TaskAdder TaskLister
type TaskAdder interface {
	AddTask(ctx context.Context, db database.Execer, t *entity.Task) error
}

type TaskLister interface {
	ListTasks(ctx context.Context, db database.Queryer) (entity.Tasks, error)
}
