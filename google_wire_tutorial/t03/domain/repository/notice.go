package repository

import "t03/domain/model"

type Notice interface {
	CreateNotice(m model.Notice) (model.Notice, error)
	GetNotice(m model.Notice) (model.Notice, error)
	ListNotice(c model.NoticeCondition) ([]model.Notice, error)
	UpdateNotice(m model.Notice) (model.Notice, error)
	DeleteNotice(m model.Notice) (model.Notice, error)
}
