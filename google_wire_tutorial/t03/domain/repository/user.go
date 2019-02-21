package repository

import (
	"context"
	"t03/domain/model"
)

type User interface {
	CreateUser(ctx context.Context, m model.User) (model.User, error)
	GetUser(ctx context.Context, m model.User) (model.User, error)
	ListUser(ctx context.Context, c model.UserCondition) ([]model.User, error)
	UpdateUser(ctx context.Context, m model.User) (model.User, error)
	DeleteUser(ctx context.Context, m model.User) (model.User, error)
}
