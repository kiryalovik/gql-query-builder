package parameter

import (
	"testing"
)

func TestPrefixedParameter_GetString(t *testing.T) {
	tests := map[string]parameterTest{
		"with alias": {
			input: &PrefixedParameter{
				name:     "product",
				alias:    "filter",
				typeName: "ProductFilter",
			},
			expectedString: "product: $filter",
			expectedPrefix: "$filter: ProductFilter",
		},
		"empty alias": {
			input: &PrefixedParameter{
				name:     "category",
				alias:    "",
				typeName: "CategoryFilter!",
			},
			expectedString: "category: $category",
			expectedPrefix: "$category: CategoryFilter!",
		},
	}
	parameterTester(t, tests)
}
