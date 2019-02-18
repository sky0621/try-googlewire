package entity

type Bookshelf interface {
}

type bookshelf struct {
}

func ProvideBookshelf() Bookshelf {
	return &bookshelf{}
}
