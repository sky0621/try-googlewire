package entity

type User interface {
}

type user struct {
}

func ProvideUser() User {
	return &user{}
}
