package repository

import (
	"users/api/pb"
	"users/internal/entity"
	"users/utils"

	helpers "github.com/zercle/gofiber-helpers"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(entity.User) (*entity.User, error)
	GetUser(uint32) (*entity.User, error)
	ListUsers(*pb.Pagination) ([]*entity.User, *entity.Pagination, error)
	UpdateUser(uint32, entity.User) (*entity.User, error)
	DeleteUser(uint32) error
}
type userRepository struct {
	*gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}
func (r *userRepository) CreateUser(req entity.User) (*entity.User, error) {
	if r.DB == nil {
		return nil, nil
	}
	err := r.DB.Create(&req).Error
	if err != nil {
		return nil, utils.NewErrorWithSource(err, helpers.WhereAmI())
	}
	return &req, nil
}
func (r *userRepository) GetUser(id uint32) (*entity.User, error) {
	user := entity.User{}
	err := r.DB.First(&user, id).Error
	if err != nil {
		return nil, utils.NewErrorWithSource(err, helpers.WhereAmI())
	}
	return &user, nil
}
func (r *userRepository) ListUsers(req *pb.Pagination) ([]*entity.User, *entity.Pagination, error) {
	users := []*entity.User{}
	err := r.DB.Limit(int(req.Limit)).Offset(int(req.Offset)).Find(&users).Error
	if err != nil {
		return nil, nil, utils.NewErrorWithSource(err, helpers.WhereAmI())
	}
	var total int64
	err = r.DB.Model(&entity.User{}).Count(&total).Error
	if err != nil {
		return nil, nil, utils.NewErrorWithSource(err, helpers.WhereAmI())
	}
	return users, &entity.Pagination{
		Total:  uint32(total),
		Limit:  req.Limit,
		Offset: req.Offset,
	}, nil
}
func (r *userRepository) UpdateUser(id uint32, req entity.User) (*entity.User, error) {
	user := entity.User{}
	err := r.DB.First(&user, id).Error
	if err != nil {
		return nil, utils.NewErrorWithSource(err, helpers.WhereAmI())
	}
	err = r.DB.Model(&user).Updates(req).Error
	if err != nil {
		return nil, utils.NewErrorWithSource(err, helpers.WhereAmI())
	}
	return &user, nil
}
func (r *userRepository) DeleteUser(id uint32) error {
	user := entity.User{}
	err := r.DB.First(&user, id).Error
	if err != nil {
		return utils.NewErrorWithSource(err, helpers.WhereAmI())
	}
	err = r.DB.Delete(&user).Error
	if err != nil {
		return utils.NewErrorWithSource(err, helpers.WhereAmI())
	}
	return nil
}
