package abuse

import "strings"

type naiveAbuseMasker struct {
	dict []string
}

func NewNaiveAbuseMasker() IAbuseMasking {
	return &naiveAbuseMasker{
		dict: []string{"fuck", "shit", "dog"},
	}
}

func (m *naiveAbuseMasker) Mask(message string) string {
	for _, word := range m.dict {
		message = strings.ReplaceAll(message, word, strings.Repeat("*", len(word)))
	}

	return message
}
