package test

import (
	"testing"

	"github.com/mittwald/revive/rule"
)

// Tests RedefinesBuiltinID rule.
func TestRedefinesBuiltinID(t *testing.T) {
	testRule(t, "redefines-builtin-id", &rule.RedefinesBuiltinIDRule{})
}
