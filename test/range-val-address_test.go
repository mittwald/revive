package test

import (
	"testing"

	"github.com/mittwald/revive/lint"
	"github.com/mittwald/revive/rule"
)

func TestRangeValAddress(t *testing.T) {
	testRule(t, "range-val-address", &rule.RangeValAddress{}, &lint.RuleConfig{})
}
