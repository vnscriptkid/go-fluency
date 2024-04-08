package persons

type person struct {
	Name string
	ID   int
}

type Manager struct {
	Title string
	person
}

func NewManager(title string, name string, id int) *Manager {
	return &Manager{
		Title: "Manager",
		person: person{
			Name: name,
			ID:   id,
		},
	}
}
