package selector

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFragment_GetFragment(t *testing.T) {
	tests := map[string]struct {
		input             *Fragment
		expectedString    string
		expectedFragments map[string]string
	}{
		"simple fragment": {
			input:          NewFragment("comparisonFields", "Character", SelectorSlice{&Field{name: "name"}}),
			expectedString: "...comparisonFields",
			expectedFragments: map[string]string{
				"comparisonFields": "fragment comparisonFields on Character{name}",
			},
		},
		"fragment in subselection": {
			input: NewFragment("comparisonFields", "Character",
				SelectorSlice{
					&Field{
						name: "name",
						subselection: SelectorSlice{
							NewFragment("comparisonFields2", "Character2", SelectorSlice{
								&Field{name: "id"},
								&Field{name: "itemCode"},
							}),
							&Field{name: "name"},
						},
					},
				}),
			expectedString: "...comparisonFields",
			expectedFragments: map[string]string{
				"comparisonFields":  "fragment comparisonFields on Character{name{...comparisonFields2 name}}",
				"comparisonFields2": "fragment comparisonFields2 on Character2{id itemCode}",
			},
		},
	}

	for testName, testCase := range tests {
		fmt.Printf("run test case: %v \n", testName)
		actualString := testCase.input.Stringify()
		assert.Exactly(t, testCase.expectedString, actualString)

		actualFragments := testCase.input.GetFragments()
		assert.Exactly(t, testCase.expectedFragments, actualFragments)
	}
}
