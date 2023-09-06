package test

import (
	"testing"

	"github.com/mittwald/revive/lint"
	"github.com/mittwald/revive/rule"
)

func TestLineLengthLimit(t *testing.T) {
	testRule(t, "line-length-limit", &rule.LineLengthLimitRule{}, &lint.RuleConfig{
		Arguments: []interface{}{int64(100)},
	})
}
