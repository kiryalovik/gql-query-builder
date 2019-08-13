package parameter

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type parameterTest struct {
	input          Parameter
	expectedString string
	expectedPrefix string
}

func parameterTester(t *testing.T, tests map[string]parameterTest) {
	for testName, testCase := range tests {
		fmt.Printf("run test case: %v \n", testName)
		actualString := testCase.input.GetString()
		assert.Exactly(t, testCase.expectedString, actualString)

		actualPrefix := testCase.input.GetPrefix()
		assert.Exactly(t, testCase.expectedPrefix, actualPrefix)
	}
}
