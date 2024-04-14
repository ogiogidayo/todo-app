package services

import (
	"context"
	"fmt"

	"github.com/ogiogidayo/todo-app/auth"

	"github.com/ogiogidayo/todo-app/database"
	"github.com/ogiogidayo/todo-app/domain"
)

type ListTask struct {
	DB   database.Queryer
	Repo TaskLister
}

func (l *ListTask) ListTasks(ctx context.Context) (domain.Tasks, error) {
	id, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user_id not found")
	}
	ts, err := l.Repo.ListTasks(ctx, l.DB, id)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return ts, err
}
