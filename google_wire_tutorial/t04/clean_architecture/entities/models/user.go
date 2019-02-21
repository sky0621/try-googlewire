package models

type User struct {
	ID string
	Name string
	Age  int
	Sex  Sex
}

type Sex int

const (
	Male Sex = iota + 1
	Female
)

type