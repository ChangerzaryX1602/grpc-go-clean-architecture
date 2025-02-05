package server

import (
	"context"
	"errors"
	"users/api/pb"
	"users/internal/usecase"
	"users/utils"

	helpers "github.com/zercle/gofiber-helpers"
)

type UserServer struct {
	usecase usecase.UserUsecase
	pb.UnimplementedUserServiceServer
}

func NewUserServer(usecase usecase.UserUsecase) *UserServer {
	return &UserServer{usecase: usecase}
}
func (s *UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	response, err := s.usecase.CreateUser(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (s *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	response, err := s.usecase.GetUser(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (s *UserServer) ListUsers(ctx context.Context, req *pb.Pagination) (*pb.ListUsersResponse, error) {
	if req.GetLimit() > 200 {
		return nil, utils.NewErrorWithSource(errors.New("limit must be less than 200"), helpers.WhereAmI())
	}
	response, err := s.usecase.ListUsers(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (s *UserServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	response, err := s.usecase.UpdateUser(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (s *UserServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	response, err := s.usecase.DeleteUser(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}
