package users

type User struct {
	Name  string
	Email string

	isMarried bool
}

func (u *User) IsMarried() bool {
	return u.isMarried
}

func New() *User {
	return &User{Name: "John Doe", Email: "john@gmail.com", isMarried: false}
}
