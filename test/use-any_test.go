package test

import (
	"testing"

	"github.com/mittwald/revive/rule"
)

func TestUseAny(t *testing.T) {
	testRule(t, "use-any", &rule.UseAnyRule{})
}
