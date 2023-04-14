package Trie

import (
	"strconv"
	"testing"
)

func Test_implTrie(t *testing.T) {
	testcases := []struct {
		name           string
		inputCommands  []string
		inputData      []string
		outputExpected []string
	}{
		{
			name:           "TC1",
			inputCommands:  []string{"Trie", "Insert", "Search", "Search", "SearchPrefix", "Insert", "Search"},
			inputData:      []string{"", "apple", "apple", "app", "app", "app", "app"},
			outputExpected: []string{"nil", "nil", "true", "false", "true", "nil", "true"},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			var trie *Trie

			for i, cmd := range tc.inputCommands {
				switch cmd {
				case "Trie":
					trie = NewTrie()
				case "Insert":
					trie.Insert(tc.inputData[i])
				case "Search":
					input := tc.inputData[i]

					actual := trie.Search(input)

					expected, _ := strconv.ParseBool(tc.outputExpected[i])

					if actual != expected {
						t.Errorf("%v [Search(%v)] Expected %v but got %v", i+1, input, expected, actual)
					}
				case "SearchPrefix":
					input := tc.inputData[i]

					actual := trie.SearchPrefix(input)

					expected, _ := strconv.ParseBool(tc.outputExpected[i])

					if actual != expected {
						t.Errorf("%v [SearchPrefix(%v)] Expected %v but got %v", i+1, input, expected, actual)
					}
				}

			}
		})
	}
}
