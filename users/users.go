package users

import (
	"program/model"
	"program/storage"
	"time"
)

type UserServer struct {
	storage storage.UserStorage
}

func NewUserServer(storage storage.UserStorage) *UserServer {
	s := &UserServer{
		storage: storage,
	}
	return s
}
func (s *UserServer) CreateUser(user model.User) error {

	founduser, err := s.storage.AssignID()
	if err != nil {
		return err
	}

	user.User_id = founduser.User_id + 1
	user.Create_time = time.Now()
	err = s.storage.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}
