package animals

type animal struct {
	Name string
	Age  int
}

func New() animal {
	return animal{Name: "Dog", Age: 3}
}
