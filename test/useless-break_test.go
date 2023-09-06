package test

import (
	"testing"

	"github.com/mittwald/revive/rule"
)

// UselessBreak rule.
func TestUselessBreak(t *testing.T) {
	testRule(t, "useless-break", &rule.UselessBreak{})
}
