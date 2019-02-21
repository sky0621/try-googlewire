package model

type UserModel interface {
}

type userModel struct {
}

func NewUser() UserModel {
	return &userModel{}
}
