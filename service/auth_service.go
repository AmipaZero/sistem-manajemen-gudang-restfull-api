package service

import (
	"errors"
	"sistem-manajemen-gudang/repository"
	"sistem-manajemen-gudang/util"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	// Register(req *RegisterRequest) error
	Login(username, password string) (string, error)
	Logout(userID uint) error
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(r repository.AuthRepository) AuthService {
	return &authService{r}
}

// type RegisterRequest struct {
// 	Nama     string     `json:"nama"`
// 	Username string     `json:"username"`
// 	Password string     `json:"password"`
// 	Role     model.Role `json:"role"` // admin / staff
// }

// func (s *authService) Register(req *RegisterRequest) error {
// 	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
// 	user := model.User{
// 		Username:     req.Username,
// 		PasswordHash: string(hash),
// 		Role:         req.Role,
// 	}
// 	return s.repo.Save(&user)
// }

func (s *authService) Login(username, password string) (string, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return "", errors.New("username not found")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", errors.New("wrong password")
	}

	token, err := util.GenerateToken(user.ID, string(user.Role))
	if err != nil {
		return "", err
	}
	s.repo.UpdateToken(user.ID, &token)
	return token, nil
}

func (s *authService) Logout(userID uint) error {
	return s.repo.UpdateToken(userID, nil)
}
