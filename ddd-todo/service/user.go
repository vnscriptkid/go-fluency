package service

import (
	domain "ddd/domain"
	repo "ddd/repo"
	"errors"
)

type UserService interface {
	CreateUser(username string) (*domain.User, error)
	GetUserByID(id int) (*domain.User, error)
}

type userService struct {
	userRepo repo.UserRepo
}

func NewUserService(userRepo repo.UserRepo) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) CreateUser(username string) (*domain.User, error) {
	if username == "" {
		return nil, errors.New("invalid username")
	}

	user, err := s.userRepo.Save(&domain.User{
		Username: username,
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetUserByID(ID int) (*domain.User, error) {
	if ID == 0 {
		return nil, errors.New("invalid ID")
	}

	user, err := s.userRepo.GetByID(ID)

	if err != nil {
		return nil, err
	}

	return user, nil
}
