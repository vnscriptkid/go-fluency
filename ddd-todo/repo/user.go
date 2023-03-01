package repo

import (
	domain "ddd/domain"
	"errors"
	"sync"
)

type UserRepo interface {
	Save(user *domain.User) (*domain.User, error)
	GetByID(id int) (*domain.User, error)
}

type UserRepoInMem struct {
	users []*domain.User
	mu    sync.Mutex
}

func (repo *UserRepoInMem) Save(user *domain.User) (*domain.User, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	user.ID = len(repo.users) + 1

	repo.users = append(repo.users, user)

	return user, nil
}

func (repo *UserRepoInMem) GetByID(id int) (*domain.User, error) {
	for _, user := range repo.users {
		if id == user.ID {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}
