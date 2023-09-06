package test

import (
	"testing"

	"github.com/mittwald/revive/rule"
)

func TestBareReturn(t *testing.T) {
	testRule(t, "bare-return", &rule.BareReturnRule{})
}
