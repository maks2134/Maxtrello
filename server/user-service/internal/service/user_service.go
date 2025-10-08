package service

import (
	"errors"
	"time"
	"user-service/internal/models"
	"user-service/internal/repository"
	"user-service/pkg/utils"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserExists         = errors.New("user already exists")
	ErrUserNotFound       = errors.New("user not found")
)

type UserService interface {
	Register(req *models.RegisterRequest) (*models.User, error)
	Login(req *models.LoginRequest) (*models.LoginResponse, error)
	GetUser(id string) (*models.User, error)
}

type userService struct {
	userRepo  repository.UserRepository
	jwtSecret string
}

func NewUserService(userRepo repository.UserRepository, jwtSecret string) UserService {
	return &userService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (s *userService) Register(req *models.RegisterRequest) (*models.User, error) {
	existingUser, err := s.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, ErrUserExists
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:    req.Email,
		Password: hashedPassword,
		Name:     req.Name,
	}

	if err := s.userRepo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) generateJWT(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}
