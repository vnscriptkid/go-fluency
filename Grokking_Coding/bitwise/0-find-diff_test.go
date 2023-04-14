package Bitwise

import "testing"

func Test_extraCharacterIndex(t *testing.T) {
	testcases := []struct {
		Name           string
		InputStr1      string
		InputStr2      string
		OutputExpected int
	}{
		{
			Name:           "TC1",
			InputStr1:      "hello",
			InputStr2:      "helo",
			OutputExpected: 2,
		},
		{
			Name:           "TC2",
			InputStr1:      "cbda",
			InputStr2:      "abc",
			OutputExpected: 2,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			OutputActual := extraCharacterIndex(tc.InputStr1, tc.InputStr2)

			if OutputActual != tc.OutputExpected {
				t.Errorf("expected %v, got %v", tc.OutputExpected, OutputActual)
			}
		})
	}
}
