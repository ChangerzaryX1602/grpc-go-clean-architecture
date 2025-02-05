package usecase

import (
	"users/api/pb"
	"users/internal/entity"
	"users/internal/repository"
)

type UserUsecase interface {
	CreateUser(*pb.CreateUserRequest) (*pb.CreateUserResponse, error)
	GetUser(*pb.GetUserRequest) (*pb.GetUserResponse, error)
	ListUsers(*pb.Pagination) (*pb.ListUsersResponse, error)
	UpdateUser(*pb.UpdateUserRequest) (*pb.UpdateUserResponse, error)
	DeleteUser(*pb.DeleteUserRequest) (*pb.DeleteUserResponse, error)
}
type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo}
}
func (u *userUsecase) CreateUser(req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := entity.User{
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}
	createdUser, err := u.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{
		Id:    uint32(createdUser.ID),
		Name:  createdUser.Name,
		Email: createdUser.Email,
	}, nil
}
func (u *userUsecase) GetUser(req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := u.repo.GetUser(req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.GetUserResponse{
		Id:    uint32(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
func (u *userUsecase) ListUsers(req *pb.Pagination) (*pb.ListUsersResponse, error) {
	users, pagination, err := u.repo.ListUsers(req)
	if err != nil {
		return nil, err
	}
	var pbUsers []*pb.User
	for _, user := range users {
		pbUsers = append(pbUsers, &pb.User{
			Id:    uint32(user.ID),
			Name:  user.Name,
			Email: user.Email,
		})
	}
	return &pb.ListUsersResponse{
		Users: pbUsers,
		Pagination: &pb.Pagination{
			Total:  uint32(pagination.Total),
			Limit:  pagination.Limit,
			Offset: pagination.Offset,
		},
	}, nil
}
func (u *userUsecase) UpdateUser(req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	user := entity.User{
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}
	updatedUser, err := u.repo.UpdateUser(req.GetId(), user)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserResponse{
		Id:    uint32(updatedUser.ID),
		Name:  updatedUser.Name,
		Email: updatedUser.Email,
	}, nil
}
func (u *userUsecase) DeleteUser(req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	err := u.repo.DeleteUser(req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.DeleteUserResponse{Success: true}, nil
}
