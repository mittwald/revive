package test

import (
	"testing"

	"github.com/mittwald/revive/lint"
	"github.com/mittwald/revive/rule"
)

func TestCommentSpacings(t *testing.T) {

	testRule(t, "comment-spacings", &rule.CommentSpacingsRule{}, &lint.RuleConfig{
		Arguments: []interface{}{"myOwnDirective"}},
	)
}
