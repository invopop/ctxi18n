package i18n_test

import (
	"context"
	"testing"

	"github.com/invopop/ctxi18n/i18n"
	"github.com/stretchr/testify/assert"
)

func TestMReplace(t *testing.T) {
	m := i18n.M{
		"string": "value",
		"num":    42,
	}
	out := m.Replace("This is a %{string} and a %{num}.")
	assert.Equal(t, "This is a value and a 42.", out)
}

func TestT(t *testing.T) {
	d := i18n.NewDict()
	d.Add("key", "value")
	l := i18n.NewLocale("en", d)
	ctx := l.WithContext(context.Background())
	assert.Equal(t, "value", i18n.T(ctx, "key"))
}

func TestN(t *testing.T) {
	d := i18n.NewDict()
	d.Add("key", map[string]any{
		"zero":  "no mice",
		"one":   "%{count} mouse",
		"other": "%{count} mice",
	})
	l := i18n.NewLocale("en", d)
	ctx := l.WithContext(context.Background())

	assert.Equal(t, "no mice", i18n.N(ctx, "key", 0, i18n.M{"count": 0}))
	assert.Equal(t, "1 mouse", i18n.N(ctx, "key", 1, i18n.M{"count": 1}))
	assert.Equal(t, "2 mice", i18n.N(ctx, "key", 2, i18n.M{"count": 2}))
}
