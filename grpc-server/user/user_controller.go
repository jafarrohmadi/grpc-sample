package user

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	protoData "grpc-server/proto"
)

type UserControllerServer interface {
	GetUserList(ctx context.Context, in *empty.Empty) (*protoData.UserList, error)
	GetUserById(ctx context.Context, in *protoData.UserId) (*protoData.User, error)
	InsertUser(ctx context.Context, in *protoData.User) (*empty.Empty, error)
	UpdateUser(ctx context.Context, in *protoData.UserUpdate) (*empty.Empty, error)
	DeleteUser(ctx context.Context, in *protoData.UserId) (*empty.Empty, error)
}
