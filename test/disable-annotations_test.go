package test

import (
	"testing"

	"github.com/mittwald/revive/lint"
	"github.com/mittwald/revive/rule"
)

func TestDisabledAnnotations(t *testing.T) {
	testRule(t, "disable-annotations", &rule.ExportedRule{}, &lint.RuleConfig{})
}

func TestModifiedAnnotations(t *testing.T) {
	testRule(t, "disable-annotations2", &rule.VarNamingRule{}, &lint.RuleConfig{})
}

func TestDisableNextLineAnnotations(t *testing.T) {
	testRule(t, "disable-annotations3", &rule.VarNamingRule{}, &lint.RuleConfig{})
}
