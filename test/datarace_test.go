package test

import (
	"testing"

	"github.com/mittwald/revive/rule"
)

func TestDatarace(t *testing.T) {
	testRule(t, "datarace", &rule.DataRaceRule{})
}
