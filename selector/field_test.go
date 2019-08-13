package selector

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gitlab.shop.by/beseller-go/gqlbuilder/parameter"
	"testing"
)

func TestField_GetString(t *testing.T) {
	tests := map[string]struct {
		input            Selector
		expectedString   string
		expectedPrefixes []string
	}{
		"field with alias no subselection and no parameters": {
			input: &Field{
				alias:  "alias",
				name:   "product",
				params: nil,
			},
			expectedString: "alias: product",
		},
		"field without alias, but with 1 level subselection and valued paramater": {
			input: &Field{
				name: "category",
				params: parameter.Container{
					parameter.NewValued("limit", 1),
				},
				subselection: []Selector{
					&Field{
						name: "id",
					},
					&Field{
						name: "itemCode",
					},
				},
			},
			expectedString: "category(limit: 1){id itemCode}",
		},
		"two leveled subselection": {
			input: &Field{
				name: "field",
				subselection: []Selector{
					&Field{
						name: "value",
						subselection: []Selector{
							&Field{
								name: "id",
							},
							&Field{name: "intValue"},
						},
					},
					&Field{
						name: "suffix",
					},
				},
			},
			expectedString: "field{value{id intValue} suffix}",
		},
		"subselection with prefixed parameters": {
			input: &Field{
				name: "field",
				params: parameter.Container{
					parameter.NewPrefixed("limit", "", "Int!"),
				},
				subselection: []Selector{
					&Field{
						name: "value",
						subselection: []Selector{
							&Field{
								name: "id",
							},
							&Field{
								name: "intValue",
								params: parameter.Container{
									parameter.NewPrefixed("filter", "find", "ValueFilter"),
								},
							},
						},
					},
					&Field{
						name: "suffix",
					},
				},
			},
			expectedString: "field(limit: $limit){value{id intValue(filter: $find)} suffix}",
			expectedPrefixes: []string{
				"$limit: Int!",
				"$find: ValueFilter",
			},
		},
	}

	for testName, testCase := range tests {
		fmt.Printf("run test case %v \n", testName)

		actualString := testCase.input.Stringify()
		assert.Exactly(t, testCase.expectedString, actualString)

		actualPrefixes := testCase.input.GetPrefixes()
		assert.Exactly(t, testCase.expectedPrefixes, actualPrefixes)
	}
}
