package test

import (
	"testing"

	"github.com/mittwald/revive/lint"
	"github.com/mittwald/revive/rule"
)

func TestCognitiveComplexity(t *testing.T) {
	testRule(t, "cognitive-complexity", &rule.CognitiveComplexityRule{}, &lint.RuleConfig{
		Arguments: []interface{}{int64(0)},
	})
}
