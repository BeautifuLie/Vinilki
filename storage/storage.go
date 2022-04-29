package storage

import "program/model"

type UserStorage interface {
	CreateUser(user model.User) error
	AssignID() (model.User, error)
}
