package service

import (
	"errors"
	"sistem-manajemen-gudang/model"
	"sistem-manajemen-gudang/repository"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)
type UserService interface {
	Register(req *RegisterRequest) error
	CurrentUser(userID uint) (*UserResponse, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{r}
}

type RegisterRequest struct {
	Username string     `json:"username"`
	Password string     `json:"password"`
	Role     model.Role `json:"role"` // admin atau staff
}
// register
func (s *userService) Register(req *RegisterRequest) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	user := model.User{
		Username:     req.Username,
		Password: string(hash),
		Role:         req.Role,
	}
	return s.repo.Save(&user)
}

type UserResponse struct {
	Username string     `json:"username"`
	Role     model.Role `json:"role"`
}
// user current
func (s *userService) CurrentUser(userID uint) (*UserResponse, error) {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user tidak ditemukan")
		}
		return nil, err
	}

	return &UserResponse{
		Username: user.Username,
		Role:     user.Role,
	}, nil
}