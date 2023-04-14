package HashMap

import (
	"strconv"
	"testing"
)

func Test_rateLimit(t *testing.T) {
	testcases := []struct {
		name           string
		inputData      [][2]string
		inputLimit     int
		outputExpected []bool
	}{
		{
			name: "TC1",
			inputData: [][2]string{
				{"2", "okay"},
				{"8", "good"},
				{"15", "okay"},
				{"16", "okay"},
				{"17", "good"},
				{"20", "okay"},
			},
			inputLimit: 5,
			outputExpected: []bool{
				true,
				true,
				true,
				false,
				true,
				false,
			},
		},
	}

	for _, tc := range testcases {
		logger := RequestLogger{}
		logger.requestLoggerInit(tc.inputLimit)

		for i, data := range tc.inputData {
			time, _ := strconv.Atoi(data[0])
			key := data[1]

			actual := logger.messageRequestDecision(time, key)
			expect := tc.outputExpected[i]

			if actual != expect {
				t.Errorf("%v. Expected %v but got %v", i+1, expect, actual)
			}
		}
	}
}
