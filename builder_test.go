package QueryBuilder

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gitlab.shop.by/beseller-go/gqlbuilder/parameter"
	"gitlab.shop.by/beseller-go/gqlbuilder/selector"
	"testing"
)

func TestBuilder_Stringify(t *testing.T) {
	tests := map[string]struct {
		input    *Builder
		expected string
		err      error
	}{
		"simple mutation": {
			input: new(Builder).
				Mutation().
				WithName("AddProduct").
				Select(
					selector.NewField("addProduct").
						WithParams(parameter.NewPrefixed("input", "products", "[ProductInput]")),
				),
			expected: "mutation AddProduct($products: [ProductInput]){addProduct(input: $products)} ",
		},
		"query with fragment": {
			input: new(Builder).Query().
				Select(
					selector.NewField("filterProduct").
						WithParams(
							parameter.NewPrefixed("filter", "", "ProductFilter"),
						).
						WithSubselection(
							selector.NewFragment("productFields", "Product", nil).
								WithSelection(
									selector.NewField("id"),
									selector.NewField("name"),
									selector.NewField("itemCode"),
								),
						),
				),
			expected: "query ($filter: ProductFilter){filterProduct(filter: $filter){...productFields}} " +
				"fragment productFields on Product{id name itemCode}",
		},
	}

	for testName, testCase := range tests {
		fmt.Printf("run test case %v \n", testName)
		actual, _ := testCase.input.Stringify()
		assert.Exactly(t, testCase.expected, actual)
	}
}
