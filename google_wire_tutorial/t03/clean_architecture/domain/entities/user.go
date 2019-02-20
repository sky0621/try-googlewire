package entities

import "clean_architecture/models"

type UserEntity interface {
	RegisterNewUser(models.User) (models.User, error)
	DeleteUsers(condition models.User) (userIDs []string, err error)
}
