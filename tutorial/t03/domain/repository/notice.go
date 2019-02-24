package repository

import (
	"context"
	"t03/domain/model"
)

type Notice interface {
	CreateNotice(ctx context.Context, m model.Notice) (model.Notice, error)
	GetNotice(ctx context.Context, m model.Notice) (model.Notice, error)
	ListNotice(ctx context.Context, c model.NoticeCondition) ([]model.Notice, error)
	UpdateNotice(ctx context.Context, m model.Notice) (model.Notice, error)
	DeleteNotice(ctx context.Context, m model.Notice) (model.Notice, error)
}
