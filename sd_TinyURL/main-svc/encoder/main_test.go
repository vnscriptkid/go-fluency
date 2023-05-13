package encoder

import "testing"

func Test_Encode(t *testing.T) {
	testcases := []struct {
		Name           string
		Input          int64
		OutputExpected string
	}{
		{
			Name:           "Test 1",
			Input:          280477000000002,
			OutputExpected: "1q3z4",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			outputActual := Encode(tc.Input)

			if outputActual != tc.OutputExpected {
				t.Errorf("Expected: %s, Got: %s", tc.OutputExpected, outputActual)
			}
		})
	}
}
