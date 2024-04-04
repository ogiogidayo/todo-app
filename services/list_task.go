package services

import (
	"context"
	"fmt"

	"github.com/ogiogidayo/todo-app/database"
	"github.com/ogiogidayo/todo-app/entity"
)

type ListTask struct {
	DB   database.Queryer
	Repo TaskLister
}

func (l *ListTask) ListTasks(ctx context.Context) (entity.Tasks, error) {
	ts, err := l.Repo.ListTasks(ctx, l.DB)
	if err != nil {
		return nil, fmt.Errorf("faild to list: %w", err)
	}
	return ts, nil
}
