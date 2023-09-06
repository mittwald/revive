package test

import (
	"testing"

	"github.com/mittwald/revive/rule"
)

func TestDuplicatedImports(t *testing.T) {
	testRule(t, "duplicated-imports", &rule.DuplicatedImportsRule{})
}
