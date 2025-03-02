package service

import (
	"sample-health/internal/model"
	"sample-health/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(user *model.User) error {
	return s.userRepo.CreateUser(user)
}

func (s *UserService) GetUserByID(id int) (*model.User, error) {
	return s.userRepo.GetUserByID(id)
}
