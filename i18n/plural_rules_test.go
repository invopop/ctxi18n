package i18n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeroOneOtherRule(t *testing.T) {
	d := &Dict{
		entries: map[string]*Dict{
			"zero":  {value: "no mice"},
			"one":   {value: "%{count} mouse"},
			"other": {value: "%{count} mice"},
		},
	}
	rule := GetRule(DefaultRuleKey)
	assert.NotNil(t, rule)
	assert.Equal(t, "no mice", rule(d, 0))
	assert.Equal(t, "%{count} mouse", rule(d, 1))
	assert.Equal(t, "%{count} mice", rule(d, 2))

	d = &Dict{
		entries: map[string]*Dict{
			"one":   {value: "%{count} mouse"},
			"other": {value: "%{count} mice"},
		},
	}
	assert.Equal(t, "%{count} mice", rule(d, 0))
	assert.Equal(t, "%{count} mouse", rule(d, 1))
	assert.Equal(t, "%{count} mice", rule(d, 2))
}
