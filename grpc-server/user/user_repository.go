package user

import protoData "grpc-server/proto"

type UserRepo interface {
	AddUser(user *protoData.User) (*protoData.User, error)
	FindUserById(id protoData.UserId) (*protoData.User, error)
	FindUsers() (*[]protoData.User, error)
	UpdateUser(user *protoData.UserUpdate) (*protoData.User, error)
	DeleteUser(id *protoData.UserId) error
}
