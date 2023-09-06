package test

import (
	"testing"

	"github.com/mittwald/revive/lint"
	"github.com/mittwald/revive/rule"
)

func TestUnusedReceiver(t *testing.T) {
	testRule(t, "unused-receiver", &rule.UnusedReceiverRule{})
	testRule(t, "unused-receiver-custom-regex", &rule.UnusedReceiverRule{}, &lint.RuleConfig{Arguments: []interface{}{
		map[string]interface{}{"allowRegex": "^xxx"},
	}})
}
