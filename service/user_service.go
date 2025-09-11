package service

import (
   "sistem-manajemen-gudang/model"
	"sistem-manajemen-gudang/repository"
	"golang.org/x/crypto/bcrypt"
)
type UserService interface {
	Register(req *RegisterRequest) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{r}
}

type RegisterRequest struct {
	Nama     string     `json:"nama"`
	Username string     `json:"username"`
	Password string     `json:"password"`
	Role     model.Role `json:"role"` // admin atau staff
}

func (s *userService) Register(req *RegisterRequest) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	user := model.User{
		Username:     req.Username,
		Password: string(hash),
		Role:         req.Role,
	}
	return s.repo.Save(&user)
}