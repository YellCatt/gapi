package service

import (
	"github.com/example/gapi/model"
	"github.com/example/gapi/repository"
)

type UserService interface {
	CreateUser(req *model.CreateUserRequest) (*model.User, error)
	GetUserByID(id uint) (*model.User, error)
	GetAllUsers() ([]model.User, error)
	UpdateUser(id uint, req *model.UpdateUserRequest) (*model.User, error)
	DeleteUser(id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(req *model.CreateUserRequest) (*model.User, error) {
	user := &model.User{
		Name: req.Name,
		Age:  req.Age,
	}
	err := s.repo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetUserByID(id uint) (*model.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) GetAllUsers() ([]model.User, error) {
	return s.repo.GetAll()
}

func (s *userService) UpdateUser(id uint, req *model.UpdateUserRequest) (*model.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Age != 0 {
		user.Age = req.Age
	}

	err = s.repo.Update(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}