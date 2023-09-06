package test

import (
	"testing"

	"github.com/mittwald/revive/lint"
	"github.com/mittwald/revive/rule"
)

func TestFunctionResultsLimit(t *testing.T) {
	testRule(t, "function-result-limit", &rule.FunctionResultsLimitRule{}, &lint.RuleConfig{
		Arguments: []interface{}{int64(3)},
	})
}
