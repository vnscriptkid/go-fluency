package domain

import "time"

type Task struct {
	ID        int
	Title     string
	Done      bool
	CreatedAt time.Time
	UserID    int
}
