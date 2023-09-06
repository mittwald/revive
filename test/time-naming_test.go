package test

import (
	"testing"

	"github.com/mittwald/revive/rule"
)

// TestTimeNamingRule rule.
func TestTimeNaming(t *testing.T) {
	testRule(t, "time-naming", &rule.TimeNamingRule{})
}
