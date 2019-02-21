package model

type User struct {
	ID   string
	Name string
	Mail string
}

type UserCondition struct {
	IDs []string
}
