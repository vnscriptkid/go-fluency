package main

import (
	"testing"
	"time"
)

func Test_getDir(t *testing.T) {
	testcases := []struct {
		name     string
		curTime  time.Time
		expected string
	}{
		{
			name:     "test1",
			curTime:  time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: "2020/01/01/00_00",
		},
		{
			name:     "test2",
			curTime:  time.Date(2020, 2, 25, 0, 0, 0, 0, time.UTC),
			expected: "2020/02/25/00_00",
		},
		{
			name:     "test3",
			curTime:  time.Date(2020, 2, 25, 0, 1, 0, 0, time.UTC),
			expected: "2020/02/25/00_00",
		},
		{
			name:     "test4",
			curTime:  time.Date(2020, 2, 25, 0, 1, 1, 0, time.UTC),
			expected: "2020/02/25/00_00",
		},
		{
			name:     "test5",
			curTime:  time.Date(2020, 2, 25, 0, 1, 1, 1, time.UTC),
			expected: "2020/02/25/00_00",
		},
		{
			name:     "test6",
			curTime:  time.Date(2020, 2, 25, 0, 4, 59, 1, time.UTC),
			expected: "2020/02/25/00_00",
		},
		{
			name:     "test7",
			curTime:  time.Date(2020, 2, 25, 5, 0, 0, 0, time.UTC),
			expected: "2020/02/25/05_00",
		},
		{
			name:     "test8",
			curTime:  time.Date(2020, 2, 25, 5, 29, 0, 0, time.UTC),
			expected: "2020/02/25/05_25",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := getDir(tc.curTime)
			if actual != tc.expected {
				t.Errorf("expected %s, got %s", tc.expected, actual)
			}
		})
	}
}

func Test_getPrevDir(t *testing.T) {
	testcases := []struct {
		name      string
		curTime   time.Time
		expected  string
		expected2 string
	}{
		{
			name:      "test1",
			curTime:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			expected:  "2019/12/31/23_55",
			expected2: "2019/12/31/23_50",
		},
		{
			name:      "test2",
			curTime:   time.Date(2020, 2, 25, 0, 0, 0, 0, time.UTC),
			expected:  "2020/02/24/23_55",
			expected2: "2020/02/24/23_50",
		},
		{
			name:      "test3",
			curTime:   time.Date(2020, 2, 25, 0, 1, 0, 0, time.UTC),
			expected:  "2020/02/24/23_55",
			expected2: "2020/02/24/23_50",
		},
		{
			name:      "test4",
			curTime:   time.Date(2020, 2, 25, 0, 1, 1, 0, time.UTC),
			expected:  "2020/02/24/23_55",
			expected2: "2020/02/24/23_50",
		},
		{
			name:      "test5",
			curTime:   time.Date(2020, 2, 25, 0, 1, 1, 1, time.UTC),
			expected:  "2020/02/24/23_55",
			expected2: "2020/02/24/23_50",
		},
		{
			name:      "test6",
			curTime:   time.Date(2020, 2, 25, 0, 4, 59, 1, time.UTC),
			expected:  "2020/02/24/23_55",
			expected2: "2020/02/24/23_50",
		},
		{
			name:      "test7",
			curTime:   time.Date(2020, 2, 25, 5, 0, 0, 0, time.UTC),
			expected:  "2020/02/25/04_55",
			expected2: "2020/02/25/04_50",
		},
		{
			name:      "test8",
			curTime:   time.Date(2020, 2, 25, 5, 29, 0, 0, time.UTC),
			expected:  "2020/02/25/05_20",
			expected2: "2020/02/25/05_15",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual1, actual2 := getPrevDir(tc.curTime)
			if actual1 != tc.expected {
				t.Errorf("expected %s, got %s", tc.expected, actual1)
			}
			if actual2 != tc.expected2 {
				t.Errorf("expected %s, got %s", tc.expected2, actual2)
			}
		})
	}
}
