package lib

import (
	"fmt"
	"os"
	"strings"
)

func ExecuteLs() (string, error) {
	wd, err := os.Getwd()

	if err != nil {
		return "", fmt.Errorf("can not get directory %w", err)
	}

	entries, err := os.ReadDir(wd)

	if err != nil {
		return "", fmt.Errorf("can not read dir %w", err)
	}

	strs := make([]string, 0)
	for _, e := range entries {
		strs = append(strs, e.Name())
	}

	return strings.Join(strs, "\t"), nil
}

// Learning points
// - wrapping error %w, useful to retain err info and easy to trace where's error from with additional context
// - os APIs: os.Getwd(), os.ReadDir()
// - looping
// - strings util: Join

func ExecutePwd() (string, error) {
	return os.Getwd()
}
