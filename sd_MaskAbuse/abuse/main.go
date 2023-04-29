package abuse

type IAbuseMasking interface {
	Mask(string) string
}
