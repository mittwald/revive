package test

import (
	"testing"

	"github.com/mittwald/revive/internal/ifelse"
	"github.com/mittwald/revive/lint"
	"github.com/mittwald/revive/rule"
)

// TestSuperfluousElse rule.
func TestSuperfluousElse(t *testing.T) {
	testRule(t, "superfluous-else", &rule.SuperfluousElseRule{})
	testRule(t, "superfluous-else-scope", &rule.SuperfluousElseRule{}, &lint.RuleConfig{Arguments: []interface{}{ifelse.PreserveScope}})
}
