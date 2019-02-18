package entity

type Book interface {
}

type book struct {
}

func ProvideBook() Book {
	return &book{}
}
