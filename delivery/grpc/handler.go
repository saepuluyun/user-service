package grpc

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"gitlab.com/zayaniqma/user-service/models"

	"golang.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_userProto "gitlab.com/zayaniqma/user-service/delivery/grpc/user"
	_user "gitlab.com/zayaniqma/user-service/user"
)

/*NewUserServer : Register User server to grpcServer
grpcServer = instance of grpc server
userService = instance of userService interface
*/
func NewUserServer(grpcServer *grpc.Server, userService _user.Service) {
	userServer := &server{
		service: userService,
	}

	_userProto.RegisterUserServiceServer(grpcServer, userServer)
	reflection.Register(grpcServer)
}

type server struct {
	service _user.Service
}

func (s *server) CreateUser(ctx context.Context, request *_userProto.CreateUserRequest) (*_userProto.CreateUserResponse, error) {
	user := s.transformCreateUserRequest(request)

	err := s.service.Store(ctx, user)

	if err != nil {
		return &_userProto.CreateUserResponse{
			Code:   "500",
			Status: "500",
			Title:  "Failed to Create User",
		}, nil
	}

	return &_userProto.CreateUserResponse{
		Code:   "201",
		Status: "201",
		Title:  "User Successfully Created",
	}, nil
}

func (s *server) transformCreateUserRequest(user *_userProto.CreateUserRequest) *models.User {
	if user == nil {
		return nil
	}

	DateOfBirth, _ := ptypes.Timestamp(user.DateOfBirth)

	res := &models.User{
		UID:         user.Uid,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
	}

	return res
}
