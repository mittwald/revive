package test

import (
	"testing"

	"github.com/mittwald/revive/lint"
	"github.com/mittwald/revive/rule"
)

func TestErrorStringsWithCustomFunctions(t *testing.T) {
	args := []interface{}{"pkgErrors.Wrap"}
	testRule(t, "error-strings-with-custom-functions", &rule.ErrorStringsRule{}, &lint.RuleConfig{
		Arguments: args,
	})
}
