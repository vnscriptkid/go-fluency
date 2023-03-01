package service

import (
	domain "ddd/domain"
	repo "ddd/repo"
	"errors"
	"time"
)

type TaskService interface {
	CreateTask(userID int, title string) (*domain.Task, error)
	GetTaskByID(ID int) (*domain.Task, error)
	GetTasksByUserID(userID int) ([]*domain.Task, error)
}

type taskService struct {
	userRepo repo.UserRepo
	taskRepo repo.TaskRepo
}

func NewTaskService(userRepo repo.UserRepo, taskRepo repo.TaskRepo) TaskService {
	return &taskService{
		userRepo: userRepo,
		taskRepo: taskRepo,
	}
}

func (s *taskService) CreateTask(userID int, title string) (*domain.Task, error) {
	if userID <= 0 {
		return nil, errors.New("invalid userID")
	}

	if title == "" {
		return nil, errors.New("invalid title")
	}

	_, err := s.userRepo.GetByID(userID)

	if err != nil {
		return nil, err
	}

	task := domain.Task{
		UserID:    userID,
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
	}

	newTask, err := s.taskRepo.Save(&task)

	if err != nil {
		return nil, err
	}

	return newTask, nil
}

func (s *taskService) GetTaskByID(ID int) (*domain.Task, error) {
	if ID == 0 {
		return nil, errors.New("invalid ID")
	}

	task, err := s.taskRepo.GetByID(ID)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *taskService) GetTasksByUserID(userID int) ([]*domain.Task, error) {
	if userID <= 0 {
		return nil, errors.New("invalid userID")
	}

	tasks, err := s.taskRepo.GetByUserID(userID)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}
