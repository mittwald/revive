package test

import (
	"testing"

	"github.com/mittwald/revive/rule"
)

func TestWaitGroupByValue(t *testing.T) {
	testRule(t, "waitgroup-by-value", &rule.WaitGroupByValueRule{})
}
