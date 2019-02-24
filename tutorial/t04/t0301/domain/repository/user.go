package repository

type User struct {
	Name string
	Age  int
	Sex  Sex
}

type Sex int

const (
	Male Sex = iota + 1
	Female
)
