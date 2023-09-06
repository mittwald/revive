package test

import (
	"testing"

	"github.com/mittwald/revive/lint"
	"github.com/mittwald/revive/rule"
)

func TestRangeValInClosure(t *testing.T) {
	testRule(t, "range-val-in-closure", &rule.RangeValInClosureRule{}, &lint.RuleConfig{})
}
