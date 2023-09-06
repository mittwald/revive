package test

import (
	"testing"

	"github.com/mittwald/revive/lint"
	"github.com/mittwald/revive/rule"
)

func TestMaxPublicStructs(t *testing.T) {
	testRule(t, "max-public-structs", &rule.MaxPublicStructsRule{}, &lint.RuleConfig{
		Arguments: []interface{}{int64(1)},
	})
}
