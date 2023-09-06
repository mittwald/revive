package test

import (
	"testing"

	"github.com/mittwald/revive/lint"
	"github.com/mittwald/revive/rule"
)

func TestContextAsArgument(t *testing.T) {
	testRule(t, "context-as-argument", &rule.ContextAsArgumentRule{}, &lint.RuleConfig{
		Arguments: []interface{}{
			map[string]interface{}{
				"allowTypesBefore": "AllowedBeforeType,AllowedBeforeStruct,*AllowedBeforePtrStruct,*testing.T",
			},
		},
	},
	)
}
