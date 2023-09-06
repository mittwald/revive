package test

import (
	"testing"

	"github.com/mittwald/revive/lint"
	"github.com/mittwald/revive/rule"
)

func TestVarNaming(t *testing.T) {
	testRule(t, "var-naming", &rule.VarNamingRule{}, &lint.RuleConfig{
		Arguments: []interface{}{[]interface{}{"ID"}, []interface{}{"VM"}},
	})

	testRule(t, "var-naming_test", &rule.VarNamingRule{}, &lint.RuleConfig{})

	testRule(t, "var-naming_upperCaseConst-false", &rule.VarNamingRule{}, &lint.RuleConfig{})
	testRule(t, "var-naming_upperCaseConst-true", &rule.VarNamingRule{}, &lint.RuleConfig{
		Arguments: []interface{}{[]interface{}{}, []interface{}{}, []interface{}{map[string]interface{}{"upperCaseConst": true}}},
	})
}
