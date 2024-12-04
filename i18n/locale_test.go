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
	assert.Equal(t, "!(MISSING: baz.random)", l.T("baz.random"))
	assert.Equal(t, "no mice", l.N("baz.mice", 0, i18n.M{"count": 0}))
	assert.Equal(t, "1 mouse", l.N("baz.mice", 1, i18n.M{"count": 1}))
	assert.Equal(t, "2 mice", l.N("baz.mice", 2, i18n.M{"count": 2}))
	assert.Equal(t, "!(MISSING: random)", l.N("random", 2))

	assert.Equal(t, "xyz", l.T("random", i18n.Default("xyz")))
	assert.Equal(t, "xyz test", l.T("random", i18n.Default("xyz %{foo}"), i18n.M{"foo": "test"}))
	assert.Equal(t, "xyz test", l.T("random", i18n.M{"foo": "test"}, i18n.Default("xyz %{foo}")))
	assert.Equal(t, "2 mouses", l.N("baz.random", 2, i18n.Default("%{count} mouses"), i18n.M{"count": 2}))
}

func TestLocaleHas(t *testing.T) {
	in := SampleLocaleData()
	l := i18n.NewLocale("en", nil)
	require.NoError(t, json.Unmarshal(in, l))

	assert.True(t, l.Has("foo"))
	assert.True(t, l.Has("baz.qux"))
	assert.False(t, l.Has("baz.random"))
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

func TestLocalUnmarshalJSON(t *testing.T) {
	l := i18n.NewLocale("en", nil)
	require.NoError(t, l.UnmarshalJSON(SampleLocaleData()))

	assert.Equal(t, "bar", l.T("foo"))

	t.Run("empty", func(t *testing.T) {
		ls := new(i18n.Locale)
		err := ls.UnmarshalJSON([]byte{})
		require.NoError(t, err)
	})
	t.Run("invalid", func(t *testing.T) {
		ls := new(i18n.Locale)
		err := ls.UnmarshalJSON([]byte("'bad'"))
		require.ErrorContains(t, err, "invalid character")
	})
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
