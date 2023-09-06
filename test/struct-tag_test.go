package test

import (
	"testing"

	"github.com/mittwald/revive/lint"
	"github.com/mittwald/revive/rule"
)

// TestStructTag tests struct-tag rule
func TestStructTag(t *testing.T) {
	testRule(t, "struct-tag", &rule.StructTagRule{})
}

func TestStructTagWithUserOptions(t *testing.T) {
	testRule(t, "struct-tag-useroptions", &rule.StructTagRule{}, &lint.RuleConfig{
		Arguments: []interface{}{"json,inline,outline", "bson,gnu"},
	})
}
