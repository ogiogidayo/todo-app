package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/ogiogidayo/todo-app/domain"
)

func (r *Repository) RegisterUser(ctx context.Context, db Execer, u *domain.User) error {
	u.Created = r.Clocker.Now()
	u.Modified = r.Clocker.Now()
	sql := `INSERT INTO user
    	(name, password, role, created, modified)
	VALUES (?, ?, ?, ?, ?)`

	result, err := db.ExecContext(ctx, sql, u.Name, u.Password, u.Role, u.Created, u.Modified)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == ErrCodeMySQLDuplicateEntry {
			return fmt.Errorf("cannot create sameuser: %w", ErrAlreadyEntry)
		}
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = domain.UserID(id)
	return nil
}

func (r *Repository) GetUser(ctx context.Context, db Queryer, name string) (*domain.User, error) {
	u := &domain.User{}
	sql := `SELECT
		id, name, password, role, created, modified
		FROM user WHERE name = ?`
	if err := db.GetContext(ctx, u, sql, name); err != nil {
		return nil, err
	}
	return u, nil
}
