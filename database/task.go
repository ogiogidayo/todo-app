package database

import (
	"context"

	"github.com/ogiogidayo/todo-app/domain"
)

func (r *Repository) ListTasks(
	ctx context.Context, db Queryer, id domain.UserID,
) (domain.Tasks, error) {
	tasks := domain.Tasks{}
	sql := `SELECT 
				id, user_id, title,
				status, created, modified 
			FROM task
			WHERE user_id = ?;`
	if err := db.SelectContext(ctx, &tasks, sql, id); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *Repository) AddTask(
	ctx context.Context, db Execer, t *domain.Task,
) error {
	t.Created = r.Clocker.Now()
	t.Modified = r.Clocker.Now()
	sql := `INSERT INTO task
			(user_id, title, status, created, modified)
	VALUES (?, ?, ?, ?, ?)`
	result, err := db.ExecContext(
		ctx, sql, t.UserID, t.Title, t.Status,
		t.Created, t.Modified,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	t.ID = domain.TaskID(id)
	return nil
}
