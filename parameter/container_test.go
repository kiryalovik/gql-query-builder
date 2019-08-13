package parameter

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainer_GetString(t *testing.T) {
	tests := map[string]struct {
		input            Container
		expectedString   string
		expectedPrefixes []string
	}{
		"only valued": {
			input: Container{
				&ValuedParameter{
					name:  "product",
					value: "some",
				},
				&ValuedParameter{
					name:  "limit",
					value: 1,
				},
			},
			expectedString:   `(product: "some", limit: 1)`,
			expectedPrefixes: nil,
		},
		"only prefixed": {
			input: Container{
				&PrefixedParameter{
					name:     "product",
					typeName: "ProductFilter",
				},
				&PrefixedParameter{
					name:     "category",
					typeName: "CategoryFilter",
					alias:    "filter",
				},
			},
			expectedString: "(product: $product, category: $filter)",
			expectedPrefixes: []string{
				"$product: ProductFilter",
				"$filter: CategoryFilter",
			},
		},
		"mixed": {
			input: Container{
				&PrefixedParameter{
					name:     "product",
					typeName: "ProductFilter",
				},
				&ValuedParameter{
					name:  "limit",
					value: 1,
				},
				&PrefixedParameter{
					name:     "category",
					typeName: "CategoryFilter",
					alias:    "filter",
				},
			},
			expectedString: "(product: $product, limit: 1, category: $filter)",
			expectedPrefixes: []string{
				"$product: ProductFilter",
				"$filter: CategoryFilter",
			},
		},
	}

	for testName, testCase := range tests {
		fmt.Printf("run test case %v \n", testName)
		actualString := testCase.input.GetString()
		assert.Exactly(t, testCase.expectedString, actualString)

		actualPrefixes := testCase.input.GetPrefixes()
		assert.Exactly(t, testCase.expectedPrefixes, actualPrefixes)
	}
}
