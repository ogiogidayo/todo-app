package services

import (
	"context"
	"fmt"

	"github.com/ogiogidayo/todo-app/database"
	"github.com/ogiogidayo/todo-app/domain"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUser struct {
	DB   database.Execer
	Repo UserRegister
}

func (r *RegisterUser) RegisterUser(
	ctx context.Context, name, password, role string,
) (*domain.User, error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("cannot hash password: %w", err)
	}
	u := &domain.User{
		Name:     name,
		Password: string(pw),
		Role:     role,
	}

	if err := r.Repo.RegisterUser(ctx, r.DB, u); err != nil {
		return nil, fmt.Errorf("failed to register: %w", err)
	}

	return u, err
}
