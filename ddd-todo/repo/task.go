package repo

import (
	domain "ddd/domain"
	"errors"
	"sync"
)

type TaskRepo interface {
	Save(user *domain.Task) (*domain.Task, error)
	GetByID(id int) (*domain.Task, error)
	GetByUserID(userID int) ([]*domain.Task, error)
}

type TaskRepoInMem struct {
	tasks []*domain.Task
	mu    sync.Mutex
}

func (repo *TaskRepoInMem) Save(task *domain.Task) (*domain.Task, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	task.ID = len(repo.tasks) + 1

	repo.tasks = append(repo.tasks, task)

	return task, nil
}

func (repo *TaskRepoInMem) GetByID(id int) (*domain.Task, error) {
	for _, task := range repo.tasks {
		if task.ID == id {
			return task, nil
		}
	}

	return nil, errors.New("task not found")
}

func (repo *TaskRepoInMem) GetByUserID(userID int) ([]*domain.Task, error) {
	list := []*domain.Task{}

	for _, task := range repo.tasks {
		if task.UserID == userID {
			list = append(list, task)
		}
	}

	return list, nil
}
