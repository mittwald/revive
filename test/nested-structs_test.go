package test

import (
	"testing"

	"github.com/mittwald/revive/lint"
	"github.com/mittwald/revive/rule"
)

func TestNestedStructs(t *testing.T) {
	testRule(t, "nested-structs", &rule.NestedStructs{}, &lint.RuleConfig{})
}
