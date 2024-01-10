package controller

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	protoData "grpc-server/proto"
	"grpc-server/user"
)

type UserController struct {
	userUsecase user.UserUsecase
	protoData.UnimplementedUsersServer
}

func NewUserController(gr *grpc.Server, userUsecase user.UserUsecase) {
	userController := &UserController{userUsecase, protoData.UnimplementedUsersServer{}}

	protoData.RegisterUsersServer(gr, userController)
}

func (e *UserController) GetUserList(ctx context.Context, in *empty.Empty) (*protoData.UserList, error) {
	users, err := e.userUsecase.FindUsers()
	if err != nil {
		return nil, err
	}
	var userx = make([]*protoData.User, 0)
	for i := 0; i < len(*users); i++ {
		var data = new(protoData.User)
		data.Id = (*users)[i].Id
		data.Email = (*users)[i].Email
		data.Name = (*users)[i].Name
		data.Address = (*users)[i].Address
		userx = append(userx, data)
	}
	var u = protoData.UserList{
		List: userx,
	}
	return &u, nil
}

func (e *UserController) GetUserById(ctx context.Context, in *protoData.UserId) (*protoData.User, error) {
	user, err := e.userUsecase.FindUserById(*in)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (e *UserController) InsertUser(ctx context.Context, in *protoData.User) (*empty.Empty, error) {
	_, err := e.userUsecase.AddUser(in)
	if err != nil {
		return new(empty.Empty), err
	}
	return new(empty.Empty), nil
}

func (e *UserController) UpdateUser(ctx context.Context, in *protoData.UserUpdate) (*empty.Empty, error) {
	_, err := e.userUsecase.UpdateUser(in)
	if err != nil {
		return new(empty.Empty), err
	}
	return new(empty.Empty), nil
}

func (e *UserController) DeleteUser(ctx context.Context, in *protoData.UserId) (*empty.Empty, error) {
	err := e.userUsecase.DeleteUser(in)
	if err != nil {
		return new(empty.Empty), err
	}
	return new(empty.Empty), nil
}
