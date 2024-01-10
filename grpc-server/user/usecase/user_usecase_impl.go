package usecase

import (
	protoData "grpc-server/proto"
	"grpc-server/user"
)

type UserUsecaseImpl struct {
	userRepo user.UserRepo
}

func NewUserUsecase(userRepo user.UserRepo) user.UserUsecase {
	return &UserUsecaseImpl{userRepo}
}

func (e *UserUsecaseImpl) AddUser(user *protoData.User) (*protoData.User, error) {
	return e.userRepo.AddUser(user)
}

func (e *UserUsecaseImpl) FindUserById(id protoData.UserId) (*protoData.User, error) {
	return e.userRepo.FindUserById(id)
}

func (e *UserUsecaseImpl) FindUsers() (*[]protoData.User, error) {
	return e.userRepo.FindUsers()
}

func (e *UserUsecaseImpl) UpdateUser(user *protoData.UserUpdate) (*protoData.User, error) {
	return e.userRepo.UpdateUser(user)
}

func (e *UserUsecaseImpl) DeleteUser(id *protoData.UserId) error {
	return e.userRepo.DeleteUser(id)
}
