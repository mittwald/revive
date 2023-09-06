package test

import (
	"testing"

	"github.com/mittwald/revive/rule"
)

// TestEmptyBlock rule.
func TestEmptyBlock(t *testing.T) {
	testRule(t, "empty-block", &rule.EmptyBlockRule{})
}
