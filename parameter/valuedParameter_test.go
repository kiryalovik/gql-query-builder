package parameter

import (
	"testing"
)

func TestValuedParameter_GetString(t *testing.T) {
	tests := map[string]parameterTest{
		"string value": {
			input: &ValuedParameter{
				name:  "product",
				value: "simple string value",
			},
			expectedString: `product: "simple string value"`,
		},
		"struct value": {
			input: &ValuedParameter{
				name: "field",
				value: struct {
					Id          int32
					StringField string
				}{
					Id:          1,
					StringField: "some string",
				},
			},
			expectedString: `field: {"Id":1,"StringField":"some string"}`,
		},
	}

	parameterTester(t, tests)
}
