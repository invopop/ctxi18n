package i18n_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/invopop/ctxi18n/i18n"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLocaleGet(t *testing.T) {
	in := SampleLocaleData()
	l := i18n.NewLocale("en", nil)

	assert.NotNil(t, l.PluralRule())
	err := json.Unmarshal(in, l)
	require.NoError(t, err)

	assert.Equal(t, "bar", l.T("foo"))
	assert.Equal(t, "quux", l.T("baz.qux"))
	assert.Equal(t, "no mice", l.N("baz.mice", 0, i18n.M{"count": 0}))
	assert.Equal(t, "1 mouse", l.N("baz.mice", 1, i18n.M{"count": 1}))
	assert.Equal(t, "2 mice", l.N("baz.mice", 2, i18n.M{"count": 2}))
	assert.Equal(t, "!(MISSING)", l.N("random", 2))
}

func TestLocaleInterpolate(t *testing.T) {
	l := i18n.NewLocale("en", nil)
	require.NoError(t, json.Unmarshal(SampleLocaleData(), l))

	out := l.N("baz.ducks", 1, 1)
	assert.Equal(t, "1 duck", out)
}

func TestLocalWithContext(t *testing.T) {
	l := i18n.NewLocale("en", nil)
	require.NoError(t, json.Unmarshal(SampleLocaleData(), l))

	l2 := i18n.GetLocale(context.Background())
	assert.Nil(t, l2)

	// Prepare Context
	ctx := l.WithContext(context.Background())
	l2 = i18n.GetLocale(ctx)
	require.NotNil(t, l2)
}

func SampleLocaleData() []byte {
	return []byte(`{
		"foo": "bar",
		"baz": {
			"qux": "quux",
			"mice": {
				"zero": "no mice",
				"one": "%{count} mouse",
				"other": "%{count} mice"
			},
			"ducks": {
				"one": "%d duck",
				"other": "%d ducks"
			}
		}
	}`)
}
