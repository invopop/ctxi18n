package i18n_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/invopop/ctxi18n/i18n"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

	ctx = context.Background()
	assert.Equal(t, "!(MISSING LOCALE)", i18n.T(ctx, "key"))
}

func TestN(t *testing.T) {
	ctx := context.Background()
	assert.Equal(t, "!(MISSING LOCALE)", i18n.N(ctx, "key", 1))

	d := i18n.NewDict()
	d.Add("key", map[string]any{
		"zero":  "no mice",
		"one":   "%{count} mouse",
		"other": "%{count} mice",
	})
	l := i18n.NewLocale("en", d)
	ctx = l.WithContext(context.Background())

	assert.Equal(t, "no mice", i18n.N(ctx, "key", 0, i18n.M{"count": 0}))
	assert.Equal(t, "1 mouse", i18n.N(ctx, "key", 1, i18n.M{"count": 1}))
	assert.Equal(t, "2 mice", i18n.N(ctx, "key", 2, i18n.M{"count": 2}))
}

func TestHas(t *testing.T) {
	d := i18n.NewDict()
	d.Add("key", "value")
	l := i18n.NewLocale("en", d)
	assert.True(t, l.Has("key"))
	assert.False(t, l.Has("missing"))
}

func TestScopes(t *testing.T) {
	in := SampleLocaleData()
	l := i18n.NewLocale("en", nil)
	require.NoError(t, json.Unmarshal(in, l))

	ctx := l.WithContext(context.Background())
	ctxScoped := i18n.WithScope(ctx, "baz")

	assert.Equal(t, "quux", i18n.T(ctxScoped, ".qux"))
	assert.True(t, i18n.Has(ctxScoped, ".qux"))
	assert.Equal(t, "!(MISSING: baz.bad)", i18n.T(ctxScoped, ".bad"))
	assert.False(t, i18n.Has(ctx, ".bad"))
	assert.Equal(t, "quux", i18n.T(ctx, "baz.qux"))
	assert.Equal(t, "!(MISSING: .qux)", i18n.T(ctx, ".qux"))
	assert.False(t, i18n.Has(ctx, ".qux"))

	assert.Equal(t, "no mice", i18n.N(ctxScoped, ".mice", 0, i18n.M{"count": 0}))

	ctxScoped = i18n.WithScope(ctxScoped, ".mice")
	assert.Equal(t, "no mice", i18n.T(ctxScoped, ".zero"))
}
