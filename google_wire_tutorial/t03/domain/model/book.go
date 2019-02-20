package model

type BookModel interface {
}

type bookModel struct {
}

func NewBook() BookModel {
	return &bookModel{}
}
