package test

import (
	"testing"

	"github.com/mittwald/revive/internal/ifelse"
	"github.com/mittwald/revive/lint"
	"github.com/mittwald/revive/rule"
)

// TestEarlyReturn tests early-return rule.
func TestEarlyReturn(t *testing.T) {
	testRule(t, "early-return", &rule.EarlyReturnRule{})
	testRule(t, "early-return-scope", &rule.EarlyReturnRule{}, &lint.RuleConfig{Arguments: []interface{}{ifelse.PreserveScope}})
}
