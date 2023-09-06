package test

import (
	"testing"

	"github.com/mittwald/revive/rule"
)

func TestGetReturn(t *testing.T) {
	testRule(t, "get-return", &rule.GetReturnRule{})
}
